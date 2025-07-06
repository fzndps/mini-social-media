document.addEventListener("DOMContentLoaded", async () => {
  const params = new URLSearchParams(window.location.search);
  let username = params.get("username");
  const token = localStorage.getItem("token");

  // Jika tidak ada di URL, pakai yang di localStorage
  if (!username) {
    username = localStorage.getItem("username");
  }

  if (!username || !token) {
    alert("Username tidak valid atau kamu belum login!");
    window.location.href = "login.html";
    return;
  }

  try {
    // Ambil data user berdasarkan username
    const resUser = await fetch(`http://localhost:3000/api/users/username/${username}`, {
      method: "GET",
      headers: { "Authorization": "Bearer " + token },
    });
    const userData = await resUser.json();
    // console.log(userData);

    if (!resUser.ok) throw new Error(userData.data || "Gagal mengambil data user");

    const user = userData.data;
    document.getElementById("profileUsername").textContent = `@${user.Username}`;

    // Ambil semua postingan berdasarkan user ID
    const resPosts = await fetch(`http://localhost:3000/api/users/profile/${user.Id}`, {
      headers: { "Authorization": "Bearer " + token },
    });
    const postData = await resPosts.json();
    // console.log(postData);

    if (!resPosts.ok) throw new Error(postData.data || "Gagal mengambil data postingan");

    const jumlahPostingan = postData.data.Posts.length;
    document.getElementById("postCount").textContent = `${jumlahPostingan} postingan`;

    const container = document.getElementById("userPosts");

    if (jumlahPostingan === 0) {
      container.innerHTML = `<p class="text-gray-500 text-center">Belum ada postingan.</p>`;
      return;
    }

    postData.data.Posts.forEach(post => {
      const div = document.createElement("div");
      div.className = "bg-white p-4 rounded shadow mb-4";
      div.innerHTML = `
        <p class="text-gray-800">${post.Content.String}</p>
        <img src="${post.ImageURL.String}" alt="Gambar"
              class="rounded object-cover max-w-full max-h-72 mx-auto block" />
        <p class="text-sm text-gray-500 mt-2">${new Date(post.CreatedAt.String).toLocaleString()}</p>
        <a href="comment.html?postId=${post.PostId.Int32}" class="text-blue-500 text-sm hover:underline">
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
      `;
      container.appendChild(div);
    });

  } catch (err) {
    console.error(err);
    alert("Terjadi kesalahan: " + err.message);
  }
});