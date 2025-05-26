import { RenderError } from "./error.js";
import { showAuthFormSignup, createBaseLayout } from "./render.js";
import { updateUserlist } from "./user.js";

// Initialize WebSocket connection
export let conn;

export function initializeWebSocket() {
  if (window["WebSocket"]) {
    conn = new WebSocket("ws://" + document.location.host + "/chat");
    conn.onopen = () => {
      conn.send(
        JSON.stringify({
          type: "message",
          message: "hi",
        })
      );
      console.log("WebSocket connection established");
    };

    conn.onmessage = (evt) => {
      try {
        const data = JSON.parse(evt.data);
        if (data.type === "all_users") {
          updateUserlist(data.users, 0);
        } else if (data.type === "Disconneted") {
          document.getElementById("user" + data.id).classList.remove("active");
        } else if (data.type === "connected") {
          if (document.getElementById("user" + data.id) != null) {
            document.getElementById("user" + data.id).classList.add("active");
          } else if (document.getElementById("user" + data.id) === null) {
            updateUserlist(data.users, 0);
          }
        } else if (data.type === "message") {
          let sender = document.getElementById("user".concat(data.sender_id));
          if (
            !sender.textContent.includes("💡") &&
            document.getElementById(`chat-${data.sender_id}`) === null
          ) {
            sender.textContent = sender.textContent + "💡";
          }
          let chat_area = document.getElementsByClassName(
            `chat-area${data.sender_id}`
          );
          if (chat_area != null) {
            AppendMessage(
              data.content,
              data.id,
              data.sender_id,
              data.sender_username
            );
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

function AppendMessage(message, id, sender_id, sender_username) {
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
  div.className = "their-message";
  div.textContent = message;
  div.prepend(time);
  box.append(div);
}
