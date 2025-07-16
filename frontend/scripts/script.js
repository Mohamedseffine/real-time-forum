import { RenderError, showNotification } from "./error.js";
import { showAuthFormSignup, createBaseLayout } from "./render.js";
import { updateUserlist } from "./user.js";

// Initialize WebSocket connection
export var conn;

export function initializeWebSocket() {
  if (window["WebSocket"]) {
    conn = new WebSocket("/chat");
    conn.onopen = () => {
      console.log("WebSocket connection established");
    };

    conn.onmessage = (evt) => {
      try {
        const data = JSON.parse(evt.data);
        if (data.type === "all_users") {
          updateUserlist(data.users, data.unreads, 0);
        } else if (data.type === "Disconneted") {
          document.getElementById("user" + data.id).classList.remove("active");
        } else if (data.type === "connected") {
          if (document.getElementById("user" + data.id) != null) {
            document.getElementById("user" + data.id).classList.add("active");
          } else if (document.getElementById("user" + data.id) === null) {
            updateUserlist(data.users, data.unreads, 0);
          }
        } else if (data.type === "message") {
          if (data.sender_id != parseInt(localStorage.getItem("id"))){
            console.log(data);
            
            showNotification(`you recieved a message from ${data.sender_username}`)
            let ulist = document.getElementsByClassName("users-list")[0];
            let sender = document.getElementById("user".concat(data.sender_id));
            ulist.prepend(sender);
            if (
              !sender.textContent.includes("💡") &&
              document.getElementById(`chat-${data.sender_id}`) === null
            ) {
              sender.textContent = sender.textContent + "💡";
            }
            let chat_area = document.getElementById(
              `chat-${data.sender_id}`
            );
            console.log(chat_area);
            
            if (chat_area != null) {
              AppendMessage(
                data.content,
                data.id,
                data.sender_id,
                data.sender_username,
              );
              let Data = {
                type: "update",
                id: data.sender_id,
                receiver_id: parseInt(localStorage.getItem("id")),
              };
              console.log("nwdfbdh:", Data);
              conn.send(JSON.stringify(Data));
            }
          }else {
            console.log("wa zaba w chta saba");
            console.log(data.reciever);
            
            let chat_area = document.getElementById(
              `chat-${data.reciever}`
            );
            console.log(chat_area);
            
            if (chat_area != null) {
              AppendMessage(
                data.content,
                data.id,
                data.reciever,
                data.sender_username,
                "sent"
              );
            }
          }
        }
      } catch (err) {
        console.error("Error parsing WebSocket message:", err);
      }
    };

    conn.onerror = (err) => {
      console.error("WebSocket error:", err);
    };

    conn.onclose = () => {
      console.log("WebSocket connection closed");
    };

    return conn;
  }
  return null;
}

export let Init = () => {
  const token = localStorage.getItem("token");
  if (location.pathname !== "/") {
    RenderError(
      "SORRY PAGE NOT FOUND",
      404,
      "The page you're looking for doesn't exist or has been moved."
    );
  } else {
    if (token) {
      createBaseLayout();
    } else {
      showAuthFormSignup();
    }
  }
};
// Initialize the application
document.addEventListener("DOMContentLoaded", () => {
  Init();
});

// Utility function for date formatting
export function formatDateFromTimestamp(ms) {
  const date = new Date(ms);
  return date.toISOString();
}

function AppendMessage(message, id, sender_id, sender_username, type="recieved") {
  const box = document.getElementById(`chat-${sender_id}`);
  if (!box) {
    return;
  }
  const div = document.createElement("div");
  const time = document.createElement("h5");
  let when = new Date(Date.now());
  when = when.toUTCString();
  time.id = "time-span";
  time.innerHTML = `${sender_username} At ${when}`;
  const br = document.createElement("div");
  time.append(br);
  div.id = "msg" + id;
  div.className =
    type === "sent"
      ? "my-message"
      : "their-message";
  div.textContent = message;
  div.prepend(time);
  box.append(div);
}
