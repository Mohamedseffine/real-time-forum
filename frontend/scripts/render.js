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

    const navBar = document.createElement('div');
    navBar.className = 'nav-bar';
    navBar.innerHTML = `
        <div class="user-info">
            <span>Welcome, ${localStorage.getItem('username') || 'User'}</span>
            <button class="logout-btn">Logout</button>
        </div>
    `;

    const container = document.createElement('div');
    container.className = 'app-container';

    const sidebar = document.createElement('div');
    sidebar.className = 'users-sidebar';
    sidebar.innerHTML = `
        <div class="users-header">
            <h2>All Users</h2>
        </div>
        <div class="users-list"></div>
    `;

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
        <div class="posts-feed"></div>
    `;

    container.appendChild(sidebar);
    container.appendChild(mainContent);
    root.appendChild(navBar);
    root.appendChild(container);

    setupPostCreation();
    setupLogoutButton();
    loadPosts();
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
            creator_id: parseInt(localStorage.getItem('id')),
            title,
            content,
            categories: selectedCategories
        };
        // if (document.cookie && document.cookie !== "") {
        //     console.log("There are cookies!");
        // } else {
        //     console.log("No cookies found.");
        // }

        try {
            const res = await fetch('/create_post', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                },
                body: JSON.stringify(payload)
            });

            if (!res.ok) throw new Error('Failed to create post');

            alert('Post created successfully!');
            document.querySelector('.post-title').value = '';
            document.querySelector('.post-creator textarea').value = '';
            document.querySelectorAll('.category.selected').forEach(el => el.classList.remove('selected'));
            loadPosts();
        } catch (err) {
            alert(`Error: ${err.message}`);
        }
    });

    document.querySelectorAll('.category').forEach(cat => {
        cat.addEventListener('click', () => {
            cat.classList.toggle('selected');
        });
    });
}

export async function loadPosts() {
    try {
        const res = await fetch('/retrieve_posts');
        const posts = await res.json();

        const feed = document.querySelector('.posts-feed');
        feed.innerHTML = '';

        for (const post of posts) {
            const postEl = document.createElement('div');
            postEl.className = 'post-item';
            postEl.innerHTML = `
                <h3 class="post-title">${post.title}</h3>
                <h5 class="post-title">${post.categorie}</h5>
                <p class="post-content">${post.content}</p>
                <div class="post-meta">
                    <span>By <strong>${post.username}</strong></span> |
                    <span>${new Date(post.creation_time).toLocaleString()}</span>
                </div>
                <div class="comments-section" data-post-id="${post.id}">
                    <h4>Comments</h4>
                    <div class="comments-list" id="comments-list-${post.id}">
                        <p class="no-comments" id="no-comments-${post.id}">No comments yet.</p>
                    </div>
                    <form class="comment-form" data-post-id="${post.id}">
                        <input type="text" class="comment-input" placeholder="Write a comment..." required />
                        <button type="submit" class="comment-btn">Post</button>
                    </form>
                </div>
            `;

            feed.appendChild(postEl);
            setupComment(post.id);
        }
    } catch (err) {
        console.error('Error loading posts:', err);
    }
}

function setupComment(postId) {
    const form = document.querySelector(`.comment-form[data-post-id="${postId}"]`);
    const commentsList = document.getElementById(`comments-list-${postId}`);
    const noComments = document.getElementById(`no-comments-${postId}`);

    retrieve_comments(postId, commentsList, noComments);

    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        const input = form.querySelector('.comment-input');
        const commentText = input.value.trim();

        if (!commentText) return;

        try {
            const res = await fetch('/create_comment', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                },
                body: JSON.stringify({
                    post_id: postId,
                    username:localStorage.getItem('username'),
                    user_id: parseInt(localStorage.getItem('id')),
                    content: commentText
                })
            });

            if (!res.ok) throw new Error('Failed to post comment');

            input.value = '';
            retrieve_comments(postId, commentsList, noComments);
        } catch (err) {
            alert('Error posting comment: ' + err.message);
        }
    });
}

function retrieve_comments(postId, commentsList, noComments) {
    fetch(`/retrieve_comments?postid=${postId}`)
        .then(res => res.json())
        .then(comments => {
            commentsList.innerHTML = '';
            if (comments.length === 0) {
                noComments.style.display = 'block';
            } else {
                noComments.style.display = 'none';
                comments.forEach(comment => {
                    const el = document.createElement('p');
                    el.textContent = `${comment.username}: ${comment.content}`;
                    commentsList.appendChild(el);
                });
            }
        })
        .catch(err => {
            console.error('Error fetching comments:', err);
        });
}
