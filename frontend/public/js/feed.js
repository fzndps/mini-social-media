const token = localStorage.getItem("token");

if (!token) {
  alert("Silakan login terlebih dahulu.");
  window.location.href = "login.html";
}

async function loadPosts() {
  try {
    const res = await fetch("http://127.0.0.1:3000/api/posts", {
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      throw new Error("Unauthorized atau gagal fetch data");
    }

    const data = await res.json();

    const feed = document.getElementById("feed");
    feed.innerHTML = "";

    data.data.forEach(post => {
      const postCard = document.createElement("div");
      postCard.className = "w-full bg-white rounded-xl shadow p-4 mb-4";

      postCard.innerHTML = `
        <div class="font-bold text-lg">${post.user.username}</div>
        <p class="text-gray-700 my-2">${post.content}</p>
        ${post.image_url ? `
        <div class="mt-2 mb-2">
          <img src="${post.image_url}" alt="Gambar"
              class="rounded object-cover max-w-full max-h-72 mx-auto block" />
        </div>` : ""}
        <div class="text-sm text-gray-500">${post.comment_count ?? 0} komentar</div>
        <a href="detail.html?postId=${post.id}" class="text-blue-500 text-sm hover:underline">Lihat Detail</a>
      `;

      feed.appendChild(postCard);
    });
  } catch (error) {
    console.error("Gagal memuat data:", error);
    document.getElementById("feed").innerHTML = `<p class="text-red-500">Gagal memuat postingan.</p>`;
  }
}


// Panggil fungsi saat halaman sudah siap
document.addEventListener("DOMContentLoaded", loadPosts);