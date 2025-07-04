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
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      throw new Error("Unauthorized atau gagal fetch data");
    }

    const data = await res.json();

    const feed = document.getElementById("feed");
    feed.innerHTML = "";

    data.data.forEach((post) => {
      const postCard = document.createElement("div");
      postCard.className = "w-full bg-white rounded-xl shadow p-4 mb-4";

      postCard.innerHTML = `
        <div class="font-bold text-lg">${post.user.username}</div>
        <p class="text-gray-700 my-2">${post.content}</p>
        ${post.image_url ? `
        <div class="mt-2 mb-2">
          <img src="${post.image_url}" alt="Gambar"
          class="rounded object-cover max-w-full max-h-72 mx-auto block" />
        </div>`
      : ""
  }
  <div class="flex justify-between items-center mt-2">  
    <div class="text-sm text-gray-500">
      ${post.comment_count ?? 0} komentar
    </div>
    <div class="flex gap-4 items-center">
      <a href="detail.html?postId=${post.id}" 
        class="inline-flex items-center text-slate-400 hover:text-slate-200 transition"
        aria-label="View Comments">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
          <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 12.75c0 1.493.876 2.783 2.25 3.422v3.328a.75.75 0 001.28.53l2.197-2.197c.312.045.63.067.956.067h7.125c2.071 0 3.75-1.679 3.75-3.75v-5.25c0-2.071-1.679-3.75-3.75-3.75H6.75c-2.071 0-3.75 1.679-3.75 3.75v5.25z"/>
        </svg>
      </a>
     
    </div>
  </div>
    `;


      feed.appendChild(postCard);
    });
  } catch (error) {
    console.error("Gagal memuat data:", error);
    document.getElementById(
      "feed"
    ).innerHTML = `<p class="text-red-500">Gagal memuat postingan.</p>`;
  }
}

// Panggil fungsi saat halaman sudah siap
document.addEventListener("DOMContentLoaded", loadPosts);
