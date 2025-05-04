if (window["WebSocket"]) {
    conn = new WebSocket("ws://" + document.location.host + "/chat")
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
            <input type="text" name="username" placeholder="Username" required>
            ${type === 'signup' ? `
            <input type="text" name="firstname" placeholder="First Name" required>
            <input type="text" name="lastname" placeholder="Last Name" required>
            <select name="gender" required>
                <option value="">Select Gender</option>
                <option value="male">Male</option>
                <option value="female">Female</option>
            </select>
            <input type="email" name="gmail" placeholder="Email" required>
            ` : ''}
            <input type="password" name="password" placeholder="Password" required>
            <button type="submit">${type === 'login' ? 'Login' : 'Sign Up'}</button>
        </form>
        <button class="back-btn">← Back</button>
    `;

    root.appendChild(formContainer);
    
    document.querySelector('.auth-form').addEventListener('submit', (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const data = Object.fromEntries(formData.entries());
        
        if (type === 'signup') {
            const { username, password, firstname, lastname, gender, gmail } = data;
            sendAuthData(gmail, username, password, firstname, lastname, gender);
        } if (type === 'login') {
            const { username, password } = data;
            sendlogindata(username, password);
        }
    });

    document.querySelector('.back-btn').addEventListener('click', () => {
        window.history.back();
    });
        // Add back button functionality
        formContainer.querySelector('.back-btn').addEventListener('click', () => {
            history.pushState(null, '', '/');
            createBaseLayout();
        });
    
}



function formatDateFromTimestamp(ms) {
    const date = new Date(ms);
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); 
    const day = String(date.getDate()).padStart(2, '0');
  
    return `${year}-${month}-${day}`;
  }
  

async function sendAuthData(email, username, password, firstname, lastname, gender) {
    let ok = Date.now();
    let createdat = formatDateFromTimestamp(ok)
    const authdata = {
        firstname: firstname,
        lastname:lastname,
        gender:gender,
        email: email,
        username: username,
        password: password,
        createdat :createdat
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
            const data = await res.json(); 
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

async function sendlogindata(username, password) {
    const type = username.includes('@') ? "email" : "username";

    const authData = {
        username: username,
        password: password,
        type: type
    };

    try {
        const res = await fetch("/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(authData)
        });

        if (res.ok) {
            const data = await res.json(); 
            console.log("Response data:", data);
            console.log("Auth data sent:", authData);
            history.pushState(null, '', '/');
            createBaseLayout();
        } else {
            console.error("Server responded with an error:", res.status);
        }
    } catch (error) {
        console.error("Error sending auth data:", error);
    }
}
