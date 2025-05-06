import { sendAuthData, sendlogindata } from "./user.js";

export function showAuthFormLogin() {
    const root = document.getElementById('root');
    root.innerHTML = '';

    const formContainer = document.createElement('div');
    formContainer.className = 'auth-form-container';

    formContainer.innerHTML = `
        <h2>Login</h2>
        <form class="auth-form">
            <input type="text" name="username" placeholder="Username" required>
            <input type="password" name="password" placeholder="Password" required>
            <button type="submit">Login</button>
        </form>
        <button class="signup-btn">Don't have an account? Sign Up</button>
    `;

    root.appendChild(formContainer);

    document.querySelector('.auth-form').addEventListener('submit', (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const data = Object.fromEntries(formData.entries());
        const { username, password } = data;
        sendlogindata(username, password);
    });

    formContainer.querySelector('.signup-btn').addEventListener('click', () => {
        showAuthFormSignup();
    });
}

export function showAuthFormSignup() {
    const root = document.getElementById('root');
    root.innerHTML = '';

    const formContainer = document.createElement('div');
    formContainer.className = 'auth-form-container';

    formContainer.innerHTML = `
        <h2>Sign Up</h2>
        <form class="auth-form">
            <input type="text" name="username" placeholder="Username" required>
            <input type="text" name="firstname" placeholder="First Name" required>
            <input type="text" name="lastname" placeholder="Last Name" required>
            <select name="gender" required>
                <option value="">Select Gender</option>
                <option value="male">Male</option>
                <option value="female">Female</option>
            </select>
            <input type="email" name="gmail" placeholder="Email" required>
            <input type="password" name="password" placeholder="Password" required>
            <button type="submit">Sign Up</button>
        </form>
        <button class="signin-btn">Already have an account? Sign In</button>
    `;

    root.appendChild(formContainer);

    document.querySelector('.auth-form').addEventListener('submit', (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const data = Object.fromEntries(formData.entries());
        const { username, password, firstname, lastname, gender, gmail } = data;
        sendAuthData(gmail, username, password, firstname, lastname, gender);
    });
    formContainer.querySelector('.signin-btn').addEventListener('click', () => {
        showAuthFormLogin();
    });
}


export function createBaseLayout() {
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

    // Main content area with updated post creator
    const mainContent = document.createElement('div');
    mainContent.className = 'main-content';
    mainContent.innerHTML = `
        <div class="post-creator">
        <input type="text" class="post-title" placeholder="Post title..." required>
        <textarea placeholder="Write your post content..."></textarea>

            <div class="category-boxes">
    <span class="category" id="1">Sport</span>
    <span class="category" id="2">Music</span>
    <span class="category" id="3">Movies</span>
    <span class="category" id="4">Science</span>
    <span class="category" id="5">Politics</span>
    <span class="category" id="6">Culture</span>
    <span class="category" id="7">Technology</span>
            </div>


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