import { createBaseLayout, showAuthFormLogin } from "./render.js";

export async function sendlogindata(username, password) {
    const type = username.includes('@') ? "email" : "username";
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
            localStorage.setItem('id', data.id);
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
            localStorage.setItem('id', data.id);
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

export async function setupLogoutButton() {
    const logoutdata  = {
            token : localStorage.getItem('token'),
            username : localStorage.getItem('username'),
            id : localStorage.getItem('id')
    }
    logoutdata.id = parseInt(logoutdata.id)
    document.querySelector('.logout-btn')?.addEventListener('click', async() => {
        try{
            const res = await fetch("/logout",{
                method : "DELETE",
                headers :{
                    "Content-Type": "application/json"
                },
                body : JSON.stringify(logoutdata)
            })
        if (res.ok){
            localStorage.removeItem('token');
            localStorage.removeItem('id');
            localStorage.removeItem('username');
            showAuthFormLogin();
        }
        }catch{
            console.error("Logout error " , error)
            alert(error.message)
        }
    });
}