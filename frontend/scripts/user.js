import { createBaseLayout, showAuthFormLogin } from "./render.js";

export async function sendlogindata(username, password) {
    const type = username.includes('@') ? "email" : "username";
    console.log(type)
    const authData = {
        username:username,
        type: type,
        password: password
    };

    try {
        const res = await fetch("/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(authData)
        });

        const data = await res.json();

        if (res.ok) {
            localStorage.setItem('token', data.token);
            localStorage.setItem('user_id', data.user_id);
            localStorage.setItem('username', data.username);
            createBaseLayout();
        } else {
            throw new Error(data.message || "Login failed");
        }
    } catch (error) {
        console.error("Login error:", error);
        alert(error.message);
    }
}

export async function sendAuthData(email, username, password, firstname, lastname, gender) {
    const authdata = {
        firstname: firstname,
        lastname: lastname,
        gender: gender,
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

        const data = await res.json();

        if (res.ok) {
            localStorage.setItem('token', data.token);
            localStorage.setItem('user_id', data.user_id);
            localStorage.setItem('username', data.username);
            createBaseLayout();
        } else {
            throw new Error(data.message || "Signup failed");
        }
    } catch (error) {
        console.error("Signup error:", error);
        alert(error.message);
    }
}

export function setupLogoutButton() {
    document.querySelector('.logout-btn')?.addEventListener('click', () => {
        localStorage.removeItem('token');
        localStorage.removeItem('user_id');
        localStorage.removeItem('username');
        showAuthFormLogin();
    });
}