import { sendAuthData, sendlogindata, setupLogoutButton } from "./user.js";

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
        <button class="switch-auth-btn">Don't have an account? Sign Up</button>
    `;

    root.appendChild(formContainer);

    document.querySelector('.auth-form').addEventListener('submit', (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const data = Object.fromEntries(formData.entries());
        sendlogindata(data.username, data.password);
    });

    formContainer.querySelector('.switch-auth-btn').addEventListener('click', showAuthFormSignup);
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
            <input type="email" name="email" placeholder="Email" required>
            <input type="password" name="password" placeholder="Password" required>
            <button type="submit">Sign Up</button>
        </form>
        <button class="switch-auth-btn">Already have an account? Sign In</button>
    `;

    root.appendChild(formContainer);

    document.querySelector('.auth-form').addEventListener('submit', (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const data = Object.fromEntries(formData.entries());
        sendAuthData(data.email, data.username, data.password, data.firstname, data.lastname, data.gender);
    });

    formContainer.querySelector('.switch-auth-btn').addEventListener('click', showAuthFormLogin);
}

export function createBaseLayout() {
    const root = document.getElementById('root');
    root.innerHTML = '';

    // Navigation bar
    const navBar = document.createElement('div');
    navBar.className = 'nav-bar';
    navBar.innerHTML = `
        <div class="user-info">
            <span>Welcome, ${localStorage.getItem('username') || 'User'}</span>
            <button class="logout-btn">Logout</button>
        </div>
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
            <!-- Users will be populated dynamically -->
        </div>
    `;

    // Main content area
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
            <!-- Posts will be loaded here -->
        </div>
    `;

    // Assemble all components
    container.appendChild(sidebar);
    container.appendChild(mainContent);
    root.appendChild(navBar);
    root.appendChild(container);

    // Add event listeners
    setupPostCreation();
    setupLogoutButton();
}

function setupPostCreation() {
    document.querySelector('.post-button')?.addEventListener('click', async () => {
        const title = document.querySelector('.post-title').value.trim();
        const content = document.querySelector('.post-creator textarea').value.trim();
        const selectedCategories = Array.from(document.querySelectorAll('.category.selected'))
            .map(el => parseInt(el.id));

        if (!title || !content) {
            alert('Please fill in both the title and content.');
            return;
        }

        const payload = {
            username: localStorage.getItem('username'),
            id: localStorage.getItem('id'),
            title: title,
            content: content,
            categories: selectedCategories

        };

        payload.id = parseInt(payload.id)
        try {
            const res = await fetch('/create_post', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                },
                body: JSON.stringify(payload)
            });
            console.log(JSON.stringify(payload));
            
            if (!res.ok) {
                throw new Error('Failed to create post');
            }

            alert('Post created successfully!');
            document.querySelector('.post-title').value = '';
            document.querySelector('.post-creator textarea').value = '';
            document.querySelectorAll('.category.selected').forEach(el => el.classList.remove('selected'));
        } catch (err) {
            alert(`Error: ${err.message}`);
        }
    });

    // Category selection
    document.querySelectorAll('.category').forEach(cat => {
        cat.addEventListener('click', () => {
            cat.classList.toggle('selected');
        });
    });
}
await setupLogoutButton()