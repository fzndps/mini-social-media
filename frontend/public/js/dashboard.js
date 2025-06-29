const postsContainer = document.getElementById("postsContainer");
const latestPostData = localStorage.getItem("latestPost");

if (latestPostData && postsContainer) {
  const post = JSON.parse(latestPostData);

  const card = document.createElement("div");
  card.className = `
    max-w-md mx-auto bg-white border border-gray-200
    rounded-2xl shadow-sm overflow-hidden mb-6
  `.trim();

  card.innerHTML = `
    <img
      src="${post.image_url}"
      class="w-full object-cover"
      alt="Post"
    />
    <div class="p-4">
      <p class="font-semibold">${post.content}</p>
    </div>
  `;

  postsContainer.prepend(card);

  // Hapus dari localStorage supaya tidak muncul lagi di reload berikutnya
  localStorage.removeItem("latestPost");
}
