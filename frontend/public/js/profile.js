document.addEventListener("DOMContentLoaded", async () => {
  const params = new URLSearchParams(window.location.search);
  const username = params.get("username");
  const token = localStorage.getItem("token");

  if (!username || !token) {
    alert("Username tidak valid atau kamu belum login!");
    window.location.href = "search.html";
    return;
  }

  try {
    // Ambil data user berdasarkan username
    const resUser = await fetch(`http://localhost:3000/api/users/username/${username}`, {
      headers: { "Authorization": "Bearer " + token },
    });
    const userData = await resUser.json();
    console.log(userData)

    if (!resUser.ok) throw new Error(userData.data || "Gagal mengambil data user");

    const user = userData.data;
    document.getElementById("profileUsername").textContent = `@${user.Username}`;

    // Ambil semua postingan berdasarkan user ID
    const resPosts = await fetch(`http://localhost:3000/api/users/profile/${user.Id}`, {
      headers: { "Authorization": "Bearer " + token },
    });
    const postData = await resPosts.json();
    console.log(postData)

    if (!resPosts.ok) throw new Error(postData.data || "Gagal mengambil data postingan");

    const container = document.getElementById("userPosts");

    if (postData.data.Posts.length === 0) {
      container.innerHTML = `<p class="text-gray-500 text-center">Belum ada postingan.</p>`;
      return;
    }

    postData.data.Posts.forEach(post => {
      const div = document.createElement("div");
      div.className = "bg-white p-4 rounded shadow";
      div.innerHTML = `
        <p class="text-gray-800">${post.Content.String}</p>
        <img src="${post.ImageURL.String}" alt="Gambar"
              class="rounded object-cover max-w-full max-h-72 mx-auto block" />
        <p class="text-sm text-gray-500 mt-2">${new Date(post.CreatedAt.String).toLocaleString()}</p>
        <a href="detail.html?postId=${post.PostId.Int32}" class="text-blue-500 text-sm hover:underline">Lihat Detail</a>
      `;
      container.appendChild(div);
    });

  } catch (err) {
    console.error(err);
    alert("Terjadi kesalahan: " + err.message);
  }
});
