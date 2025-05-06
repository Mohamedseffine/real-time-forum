import { showAuthFormSignup , createBaseLayout,showAuthFormLogin} from "./render.js"

if (window["WebSocket"]) {
    var conn 
    conn = new WebSocket("ws://" + document.location.host + "/chat")
    conn.onopen = () => {
        console.log("websockets open lol")
    }
    conn.onmessage = (evt) => {
        const data = JSON.parse(evt.data)
        console.log("received data : ", data)
    }
    conn.onerror = (err) => {
        console.log("error websockets", err);
    }
}



// Initialize the application
showAuthFormSignup();

// document.addEventListener('DOMContentLoaded', () => {
//     document.body.addEventListener('click', (e) => {
//         if (e.target.classList.contains('login-btn')) {
//             history.pushState(null, '', '/login');
//             showAuthForm('login');
//         } else if (e.target.classList.contains('signup-btn')) {
//             history.pushState(null, '', '/signup');
//             showAuthForm('signup');
//         }
//     });
// });



export function formatDateFromTimestamp(ms) {
    const date = new Date(ms);
    return date.toISOString();
}




document.addEventListener('DOMContentLoaded', () => {
    // Toggle category selection
    document.querySelectorAll('.category').forEach(cat => {
        cat.addEventListener('click', () => {
            cat.classList.toggle('selected');
        });
    });

    // Handle post submission
    document.querySelector('.post-button').addEventListener('click', async () => {
        const title = document.querySelector('.post-title').value.trim();
        const content = document.querySelector('.post-creator textarea').value.trim();
        const selectedCategories = Array.from(document.querySelectorAll('.category.selected'))
  .map(el => parseInt(el.id));


        if (!title || !content) {
            alert('Please fill in both the title and content.');
            return;
        }

        const payload = {
            username: "JohnDoe", //replace when fix
            id: 1,
            title: title,
            content: content,
            categories: selectedCategories
        };

        try {
            const res = await fetch('/create_post', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload)
            });

            if (!res.ok) {
                const errorText = await res.text(); 
                throw new Error(errorText || 'Failed to create post');
            }

            alert('Post created successfully!');
            console.log(payload);
            document.querySelector('.post-title').value = '';
            document.querySelector('.post-creator textarea').value = '';
            document.querySelectorAll('.category.selected').forEach(el => el.classList.remove('selected'));
        } catch (err) {
            alert(`Error: ${err.message}`);
        }
    });
});
