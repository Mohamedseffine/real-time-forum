import { createBaseLayout, showAuthFormLogin } from "./render.js";
import { conn } from "./script.js";
var LaStInsertedId = 0;
export async function sendlogindata(username, password) {
  const type = username.includes("@") ? "email" : "username";
  const authData = {
    username: username,
    type: type,
    password: password,
  };

  try {
    const res = await fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(authData),
    });

    const data = await res.json();

    if (res.ok) {
      localStorage.setItem("token", data.token);
      localStorage.setItem("id", data.id);
      localStorage.setItem("username", data.username);
      createBaseLayout();
    } else {
      throw new Error(data.message || "Login failed");
    }
  } catch (error) {
    console.error("Login error:", error);
    alert(error.message);
  }
}

export async function sendAuthData(
  email,
  username,
  password,
  firstname,
  lastname,
  gender
) {
  const authdata = {
    firstname: firstname,
    lastname: lastname,
    gender: gender,
    email: email,
    username: username,
    password: password,
  };

  try {
    const res = await fetch("/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(authdata),
    });

    const data = await res.json();
    if (res.ok) {
      localStorage.setItem("token", data.token);
      localStorage.setItem("id", data.id);
      localStorage.setItem("username", data.username);
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
  const logoutdata = {
    token: localStorage.getItem("token"),
    username: localStorage.getItem("username"),
    id: localStorage.getItem("id"),
  };
  logoutdata.id = parseInt(logoutdata.id);
  document.querySelector(".logout-btn")?.addEventListener("click", async () => {
    try {
      const res = await fetch("/logout", {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(logoutdata),
      });
      if (res.ok) {
        localStorage.removeItem("token");
        localStorage.removeItem("id");
        localStorage.removeItem("username");
        conn.close();
        showAuthFormLogin();
      }
    } catch {
      console.error("Logout error ", error);
      alert(error.message);
    }
  });
}
export function setupComment(postId, commentsList, noCommentsEl) {
  const form = commentsList
    .closest(".comments-section")
    .querySelector(".comment-form");

  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const input = form.querySelector(".comment-input");
    const content = input.value.trim();

    if (!content) return;

    const payload = {
      user_id: parseInt(localStorage.getItem("id")),
      postid: postId,
      username: localStorage.getItem("username"),
      content: content,
    };

    try {
      const res = await fetch("/create_comment", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify(payload),
      });

      if (!res.ok) throw new Error("Failed to post comment");
      const savedComment = await res.json();
      console.log(savedComment);

      input.value = "";
      noCommentsEl?.remove();
    } catch (err) {
      alert(`Error: ${err.message}`);
    }
  });
}
export function updateUserlist(users, id) {
  const userList = document.querySelector(".users-list");
  userList.innerHTML = "";

  if (users.length === 0) {
    userList.innerHTML = "<p>No users found.</p>";
    return;
  }

  users.forEach((user) => {
    if (user.id === parseInt(localStorage.getItem("id"))) return;
    if (user.id === id) {
      console.log("hahaha");

      return;
    }
    const userItem = document.createElement("button");
    userItem.className = "user-item";
    userItem.textContent = user.username;
    userItem.id = "user" + user.id;
    userItem.dataset.userid = user.id;

    if (user.active === 1) {
      userItem.classList.add("active");
    }

    // Attach event to open chat
    userItem.addEventListener("click", async () => {
      LaStInsertedId = 0;
      userItem.innerText = userItem.innerText.replace("💡", "");
      openChatWithUser(user);
      LaStInsertedId = await getMessages(
        parseInt(localStorage.getItem("id")),
        user.id,
        LaStInsertedId
      );
      console.log('jjjjjjjjjjjjjjjj');
      
    });

    userList.appendChild(userItem);
  });
}

// Example placeholder function to open chat (implement your chat logic here)
function openChatWithUser(user) {
  console.log("last id", LaStInsertedId);

  let chatArea = document.querySelector(".chat-area");

  // If it already exists, remove it before creating a new one
  if (chatArea) {
    chatArea.remove();
  }

  chatArea = document.createElement("div");
  chatArea.className = "chat-area";

  chatArea.innerHTML = `
        <div class="chat-header">
            Chat with ${user.username}
            <button class="close-chat-btn">✖</button>
        </div>
        <div class="chat-messages" id="chat-${user.id}">
        
        </div>
        <input type="text" class="chat-input" placeholder="Type a message...">
        <button class="send-btn">Send</button>
    `;

  document.body.appendChild(chatArea);
  
  const closeBtn = chatArea.querySelector(".close-chat-btn");
  closeBtn.addEventListener("click", () => {
    LaStInsertedId = 0;
    chatArea.remove();
  });

  const sendBtn = chatArea.querySelector(".send-btn");
  sendBtn.addEventListener("click", () => {
    sendMessage(user.id);
  });

  const cont = document.getElementById(`chat-${user.id}`);
  cont.addEventListener("scrollend", async()=>{
    const { scrollTop } = cont;

    if (scrollTop === 0) {
      console.log(LaStInsertedId);
      LaStInsertedId = await getMessages(parseInt(localStorage.getItem("id")), user.id, LaStInsertedId)
      
    }
  });
}

function sendMessage(userId) {
  const input = document.querySelector(".chat-input");
  const message = input.value.trim();
  if (!message) return;

  const messageBox = document.getElementById(`chat-${userId}`);
  const msgElement = document.createElement("div");
  msgElement.className = "my-message";
  msgElement.textContent = message;
  messageBox.appendChild(msgElement);
  let username = document
    .getElementById("user".concat(userId))
    .textContent.replace("💡", "");
  input.value = "";
  let msg = {
    type: "message",
    message: message,
    id: parseInt(localStorage.getItem("id")),
    username: localStorage.getItem("username"),
    receiver_id: userId,
    receiver_username: username,
    status: "unread",
  };
  if (conn.readyState === WebSocket.OPEN) {
    conn.send(JSON.stringify(msg));
  } else {
    console.log("websocket not open");
  }
}

export async function getMessages(senderId, receiverId, lastID) {
  if (lastID === null) {
    return null;
  }
  const payload = {
    sender_id: senderId,
    receiver_id: receiverId,
    last_id: lastID, // ask server for n most-recent messages
  };
  

  try {
    const res = await fetch("/get_chat", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      body: JSON.stringify(payload),
    });

    if (!res.ok) {
      throw new Error("Failed to fetch messages");
    }

    const messages = await res.json();

    
    const box = document.getElementById(`chat-${receiverId}`);
    console.log(box);
    
    if (!box) {
      return null;
    }
    console.log(messages["messages"]);

    if (messages["messages"] === null) {
      return null;
    }
    messages["messages"].forEach((msg) => {
      const div = document.createElement("div");
      const mine = msg.user_id === senderId;
      div.id = "msg" + msg.id;
      div.className = mine ? "my-message" : "their-message";
      div.textContent = msg.message;
      box.prepend(div);
    });

    box.scrollTop = box.scrollHeight;
    // Hand back the earliest message-id so caller can request the
    // previous page next time (useful for infinite scroll).
    return messages["messages"].length == 10
      ? messages["messages"][messages["messages"].length - 1].id
      : null;
  } catch (err) {
    console.error("getMessages error:", err);
    alert(err.message);
  }
}


