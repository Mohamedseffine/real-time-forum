import {createBaseLayout} from "./render.js"
import { formatDateFromTimestamp } from "./script.js";
export async function sendlogindata(username, password) {
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
            localStorage.setItem('token', data.token);
            localStorage.setItem('user_id', data.user_id);
            localStorage.setItem('username', data.username);
            history.pushState(null, '', '/home');
            createBaseLayout();
        } else {
            console.error("Login failed:", res.status);
        }
    } catch (error) {
        console.error("Error:", error);
    }
}


export async function sendAuthData(email, username, password, firstname, lastname, gender) {
    const createdat = formatDateFromTimestamp(Date.now());
    const authdata = {
        firstname: firstname,
        lastname: lastname,
        gender: gender,
        email: email,
        username: username,
        password: password,
        createdat: createdat
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
            console.log(data);
            
            // localStorage.setItem('token', data.token);
            // localStorage.setItem('user_id', data.user_id);
            // localStorage.setItem('username', data.username);
            history.pushState(null, '', '/home');
            createBaseLayout();
        } else {
            console.log("Server error:", res.status);
        }
    } catch (error) {
        console.log("Error:", error);
    }
}
