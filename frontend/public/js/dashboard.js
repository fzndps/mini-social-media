document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");

  if (!token) {
    alert("Silakan login terlebih dahulu.");
    window.location.href = "login.html";
    return;
  }

  loadPosts();

  async function loadPosts() {
    const feed = document.getElementById("feed");
    feed.innerHTML = "<p class='text-gray-500'>Memuat postingan...</p>";


    const newPostJson = localStorage.getItem("newPost");
    // Cek postingan baru di localStorage
    if (newPostJson && newPostJson !== "undefined") {
        const newPost = JSON.parse(newPostJson);
        const newCard = createPostCard(newPost);
        feed.innerHTML = ""; // Kosongkan feed dulu
        feed.appendChild(newCard);
        localStorage.removeItem("newPost");
    }

    
    try {
      const response = await fetch("http://127.0.0.1:3000/api/posts", {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error(`Gagal memuat data: ${response.status}`);
      }

      const data = await response.json();
      feed.innerHTML = "";

      data.data.forEach((post) => {
        const card = createPostCard(post);
        feed.appendChild(card);
      });
    } catch (error) {
      console.error("Error:", error);
      feed.innerHTML =
        `<p class="text-red-500">Gagal memuat postingan. Silakan refresh halaman.</p>`;
    }
  }

  function createPostCard(post) {
    const card = document.createElement("div");
    card.className = "w-full bg-white rounded-xl shadow p-4 mb-4 relative";

    card.innerHTML = `
      <div class="flex justify-between items-start">
        <div class="font-bold text-lg">${post.user.username}</div>
        <button class="menu-toggle">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none"
            viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor"
            class="w-6 h-6 text-slate-500 hover:text-slate-700 cursor-pointer transition">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M6.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM18.75 12a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z" />
          </svg>
        </button>
        <div class="menu-dropdown hidden absolute right-4 top-8 bg-white border rounded shadow-lg z-10">
          <button class="delete-post w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50">
            Delete Post
          </button>
        </div>
      </div>

      <p class="text-gray-700 my-2">${post.content}</p>
      ${
        post.image_url? `
          <div class="mt-2 mb-2">
            <img src="${post.image_url}" alt="Gambar"
              class="rounded object-cover max-w-full max-h-72 mx-auto block" />
          </div>`
          : ""
      }

      <div class="mt-4 flex flex-col items-start space-y-1">
        <a href="comment.html?postId=${post.id}"
          class="inline-flex items-center text-slate-400 hover:text-slate-600 transition"
          aria-label="View Comments">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none"
            viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M12 20.25c4.97 0 9-3.694 9-8.25s-4.03-8.25-9-8.25S3 7.444 3 12c0 2.104.859 4.023 2.273 5.48.432.447.74 1.04.586 1.641a4.483 4.483 0 0 1-.923 1.785A5.969 5.969 0 0 0 6 21c1.282 0 2.47-.402 3.445-1.087.81.22 1.668.337 2.555.337Z" />
          </svg>
        </a>
        <div class="text-sm text-gray-500">
          ${post.comment_count ?? 0} komentar
        </div>
      </div>
    `;

    // Toggle dropdown menu
    const menuToggle = card.querySelector(".menu-toggle");
    const menuDropdown = card.querySelector(".menu-dropdown");

    menuToggle.addEventListener("click", (e) => {
      e.stopPropagation();
      menuDropdown.classList.toggle("hidden");
    });

    document.addEventListener("click", () => {
      menuDropdown.classList.add("hidden");
    });

    // Delete handler
    const deleteButton = card.querySelector(".delete-post");
    deleteButton.addEventListener("click", async () => {
      if (confirm("Apakah kamu yakin ingin menghapus postingan ini?")) {
        try {
          const response = await fetch(`http://127.0.0.1:3000/api/posts/${post.id}/user/${post.user.id}`, {
            method: "DELETE",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${token}`,
            },
          });

          if (response.ok) {
            card.remove();
          } else {
            const errorText = await response.text();
            alert(`Gagal menghapus postingan: ${errorText}`);
          }
        } catch (error) {
          console.error("Gagal menghapus postingan:", error);
          alert("Terjadi kesalahan saat menghapus postingan.");
        }
      }
    });

    return card;
  }
});
