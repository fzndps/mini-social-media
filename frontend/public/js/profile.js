document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  let username = localStorage.getItem("username");

  if (!token) {
    alert("Silakan login terlebih dahulu.");
    window.location.href = "login.html";
    return;
  }

  loadProfile();

  async function loadProfile() {
    const profileUsername = document.getElementById("profileUsername");
    const postCount = document.getElementById("postCount");
    const userPosts = document.getElementById("userPosts");

    try {
      // Jika username belum tersimpan, cari dulu dari postingan
      if (!username) {
        const postRes = await fetch(`http://127.0.0.1:3000/api/posts`, {
          method: "GET",
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (!postRes.ok) {
          throw new Error("Gagal mengambil postingan");
        }

        const postData = await postRes.json();

        // Cari postingan milik user ini
        const firstPost = postData.data.find((post) => post.user && post.user.username);
        

        if (!firstPost) {
          userPosts.innerHTML = `<p class="text-gray-500">Tidak dapat menentukan pengguna dari postingan.</p>`;
          profileUsername.textContent = "";
          postCount.textContent = "";
          return;
        }

        username = firstPost.user.username;
        // Simpan username untuk next load
        localStorage.setItem("username", username);
      }

      // Ambil data user berdasarkan username
      const userRes = await fetch(`http://127.0.0.1:3000/api/users/username/${username}`, {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!userRes.ok) {
        throw new Error("Gagal mengambil data profil");
      }

      const userData = await userRes.json();
      const userId = userData.data.id;

      profileUsername.textContent = `@${userData.data.username}`;

      // Ambil semua postingan lagi
      const postRes = await fetch(`http://127.0.0.1:3000/api/posts`, {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!postRes.ok) {
        throw new Error("Gagal mengambil postingan");
      }

      const postData = await postRes.json();

      const userPostsData = postData.data.filter(
        (post) => post.user_id === userId || (post.user && post.user.id === userId)
      );

      postCount.textContent = `${userPostsData.length} postingan`;

      if (userPostsData.length === 0) {
        userPosts.innerHTML = `<p class="text-gray-500">Belum ada postingan.</p>`;
        return;
      }

      // Bersihkan loader
      userPosts.innerHTML = "";

      // Render postingan
      userPostsData.forEach((post) => {
        const card = document.createElement("div");
        card.className = "bg-white rounded-xl shadow p-4";
        card.innerHTML = `
          <p class="text-gray-700 mb-2">${post.content}</p>
          ${
            post.image_url
              ? `<img src="${post.image_url}" alt="Gambar" class="rounded max-h-64 object-cover mx-auto"/>`
              : ""
          }
          <div class="text-sm text-gray-500 mt-2">${post.comment_count ?? 0} komentar</div>
        `;
        userPosts.appendChild(card);
      });

    } catch (error) {
      console.error(error);
      userPosts.innerHTML = `<p class="text-red-500">${error.message}</p>`;
    }
  }
});
