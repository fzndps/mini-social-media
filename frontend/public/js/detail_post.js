const token = localStorage.getItem("token");
if (!token) {
  alert("Silakan login terlebih dahulu.");
  window.location.href = "login.html";
}

function getPostIdFromURL() {
  const params = new URLSearchParams(window.location.search);
  return params.get("postId");
}

const postId = getPostIdFromURL();

async function fetchPostDetail() {
  try {
    const res = await fetch(`http://127.0.0.1:3000/api/posts/${postId}/comments`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!res.ok) throw new Error("Gagal mengambil detail post");

    const data = await res.json();
    const post = data.data;
    const comments = data.data.Comments;

    renderPost(post);
    renderComments(comments);
  } catch (error) {
    console.error("Gagal memuat detail:", error);
  }
}

function renderPost(post) {
  const container = document.getElementById("postDetail");
  container.innerHTML = `
    <h2 class="text-xl font-bold mb-2">${post.Content}</h2>
    ${post.ImageURL ? `<img src="${post.ImageURL}" class="w-full rounded mb-2"/>` : ""}
    <p class="text-sm text-gray-500">Dipost oleh ${post.Username}</p>
  `;
}

function renderComments(comments) {
  const list = document.getElementById("commentList");
  list.innerHTML = "";

  comments.forEach((comment) => {
    const div = document.createElement("div");
    div.className = "p-3 bg-white rounded shadow";

    div.innerHTML = `
      <div>
        <p class="text-md"><strong>${comment.User.Username}</strong></p>
        <p class="text-md">${comment.Content}</p>
        <p class="text-xs text-gray-500">${new Date(comment.CreatedAt).toLocaleString()}</p>
        <div class="flex space-x-2 text-sm mt-1">
          <button onclick="editComment(${comment.Id}, '${comment.Content}')" class="text-blue-600 hover:underline">Edit</button>
          <button onclick="deleteComment(${comment.Id})" class="text-red-600 hover:underline">Hapus</button>
        </div>
      </div>
    `;
    list.appendChild(div);
  });
}

async function submitComment() {
  const content = document.getElementById("newCommentInput").value.trim();
  if (!content) return;

  await fetch(`http://127.0.0.1:3000/api/posts/${postId}/comments`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ content }),
  });

  document.getElementById("newCommentInput").value = "";
  fetchPostDetail();
}

async function editComment(commentId, oldContent) {
  const newContent = prompt("Edit komentar:", oldContent);
  if (!newContent || newContent === oldContent) return;

  await fetch(`http://127.0.0.1:3000/api/comments/${commentId}`, {
    method: "PUT",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ content: newContent }),
  });

  fetchPostDetail();
}

async function deleteComment(commentId) {
  const confirmed = confirm("Yakin ingin hapus komentar ini?");
  if (!confirmed) return;

  await fetch(`http://127.0.0.1:3000/api/posts/${postId}/comments/${commentId}`, {
    method: "DELETE",
    headers: { Authorization: `Bearer ${token}` },
  });

  fetchPostDetail();
}

document.addEventListener("DOMContentLoaded", () => {
  fetchPostDetail();
  document.getElementById("submitCommentBtn").addEventListener("click", submitComment);
});
