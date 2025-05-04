if (window["WebSocket"]) {
    conn = new WebSocket("ws://" + document.location.host + "/ws")
    conn.onopen =()=>{
        console.log("websockets open lol")
    }
    conn.onmessage =(evt) => {
        const data = JSON.parse(evt.data)
        console.log("received data : ",data)
    }
    conn.onerror =(err)=>{
        console.log("error websockets" , err);
    }
}
function createBaseLayout() {
    const root = document.getElementById('root');
    root.innerHTML = '';

    // Create authentication buttons
    const authButtons = document.createElement('div');
    authButtons.className = 'auth-buttons';
    authButtons.innerHTML = `
        <button class="login-btn">Login</button>
        <button class="signup-btn">Sign Up</button>
    `;

    // Main app container
    const container = document.createElement('div');
    container.className = 'app-container';

    // Users sidebar
    const sidebar = document.createElement('div');
    sidebar.className = 'users-sidebar';
    sidebar.innerHTML = `
        <div class="users-header">
            <h2>Active Users</h2>
        </div>
        <div class="users-list">
            <div class="user-item" data-user-id="1">
                <div class="user-avatar"></div>
                <span class="username">JohnDoe</span>
            </div>
            <div class="user-item" data-user-id="2">
                <div class="user-avatar"></div>
                <span class="username">JaneSmith</span>
            </div>
        </div>
    `;

    // Main content area
    const mainContent = document.createElement('div');
    mainContent.className = 'main-content';
    mainContent.innerHTML = `
        <div class="post-creator">
            <textarea placeholder="Create a new post..."></textarea>
            <button class="post-button">Post</button>
        </div>
        
        <div class="posts-feed">
            <!-- Sample Post 1 -->
            <div class="post">
                <div class="post-header">
                    <div class="author-avatar"></div>
                    <span class="author-name">JohnDoe</span>
                </div>
                <div class="post-content">
                    This is a sample post with some example content
                </div>
                <div class="post-actions">
                    <button class="like-button">👍 24</button>
                    <button class="comment-button">💬 5</button>
                </div>
                <div class="comments-section">
                    <div class="comment">
                        <div class="comment-author">JaneSmith</div>
                        <div class="comment-text">Great post!</div>
                        <button class="like-button">👍 3</button>
                    </div>
                    <div class="comment-input">
                        <input type="text" placeholder="Write a comment...">
                        <button>Post</button>
                    </div>
                </div>
            </div>
        </div>
    `;

    // Chat container (hidden by default)
    const chatContainer = document.createElement('div');
    chatContainer.className = 'chat-container';
    chatContainer.innerHTML = `
        <div class="chat-header">
            <h3>Chat with <span class="chat-partner">Username</span></h3>
            <button class="close-chat">×</button>
        </div>
        <div class="chat-messages">
            <div class="message incoming">
                <div class="message-content">Hello!</div>
            </div>
            <div class="message outgoing">
                <div class="message-content">Hi there!</div>
            </div>
        </div>
        <div class="chat-input">
            <input type="text" placeholder="Type your message...">
            <button>Send</button>
        </div>
    `;

    // Assemble all components
    container.appendChild(sidebar);
    container.appendChild(mainContent);
    container.appendChild(chatContainer);
    
    root.appendChild(authButtons);
    root.appendChild(container);
}

// Initialize the application
createBaseLayout();
document.addEventListener('DOMContentLoaded', () => {
    document.body.addEventListener('click', (e) => {
        if (e.target.classList.contains('login-btn')) {
            history.pushState(null, '', '/login');
            showAuthForm('login');
        } else if (e.target.classList.contains('signup-btn')) {
            history.pushState(null, '', '/signup');
            showAuthForm('signup');
        }
    });
});

function showAuthForm(type) {
    const root = document.getElementById('root');
    root.innerHTML = ''; 

    const formContainer = document.createElement('div');
    formContainer.className = 'auth-form-container';

    formContainer.innerHTML = `
        <h2>${type === 'login' ? 'Login' : 'Sign Up'}</h2>
        <form class="auth-form">
            <input type="text" name="gmail" placeholder="gmail" required>
            <input type="text" name="username" placeholder="Username" required>
            <input type="password" name="password" placeholder="Password" required>
            <button type="submit">${type === 'login' ? 'Login' : 'Sign Up'}</button>
        </form>
        <button class="back-btn">← Back</button>
    `;

    root.appendChild(formContainer);
    
document.querySelector('.auth-form').addEventListener('submit', (e) => {
    e.preventDefault();
    const username = e.target.querySelector('input[name="username"]').value;
    const password = e.target.querySelector('input[name="password"]').value;
    const email = e.target.querySelector('input[name="gmail"]').value
    const type = window.location.pathname.includes('signup') ? 'signup' : 'login';
    console.log(type);
    if (type ==  "signup"){
        sendAuthData(email, username, password);
    }else{
        sendlogindata(email , username, password)
    }
 
   
});


    // Add back button functionality
    formContainer.querySelector('.back-btn').addEventListener('click', () => {
        history.pushState(null, '', '/');
        createBaseLayout();
    });

}

async function sendAuthData(email, username, password) {
    const authdata = {
        email: email,
        username: username,
        password: password
    };

    try {
        const res = await fetch("/signup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(authdata)
        });

        if (res.ok) {
            const data = await res.json(); // or res.text() if not JSON
            console.log("Response data:", data);
            console.log("Auth data sent:", authdata);
            history.pushState(null, '', '/');
            createBaseLayout();
        } else {
            console.log("Server responded with an error:", res.status);
        }
    } catch (error) {
        console.log("Error sending auth data:", error);
    }
}

async function sendlogindata(email, username, password) {
    const authdata = {
        email: email,
        username: username,
        password: password
    };

    try {
        const res = await fetch("/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(authdata)
        });

        if (res.ok) {
            const data = await res.json(); // or res.text() if not JSON
            console.log("Response data:", data);
            console.log("Auth data sent:", authdata);
            history.pushState(null, '', '/');
            createBaseLayout();
        } else {
            console.log("Server responded with an error:", res.status);
        }
    } catch (error) {
        console.log("Error sending auth data:", error);
    }
}
