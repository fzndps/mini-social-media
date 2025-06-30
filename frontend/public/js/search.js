document.addEventListener("DOMContentLoaded", () => {
  const searchBtn = document.getElementById("searchBtn");
  const searchInput = document.getElementById("searchInput");
  const userCards = document.getElementById("userCards");
  const searchResults = document.getElementById("searchResults");
  const noResults = document.getElementById("noResults");
  const errorMessage = document.getElementById("errorMessage");
  const errorText = document.getElementById("errorText");
  const loadingIndicator = document.getElementById("loadingIndicator");

  searchBtn.addEventListener("click", async () => {
    const username = searchInput.value.trim();
    const token = localStorage.getItem("token");

    // Reset semua tampilan
    userCards.innerHTML = "";
    searchResults.classList.add("hidden");
    noResults.classList.add("hidden");
    errorMessage.classList.add("hidden");
    loadingIndicator.classList.remove("hidden");

    if (!username) {
      alert("Masukkan username yang ingin dicari!");
      loadingIndicator.classList.add("hidden");
      return;
    }

    if (!token) {
      alert("Kamu belum login!");
      window.location.href = "login.html";
      return;
    }

    try {
      const res = await fetch(`http://localhost:3000/api/users/username/${username}`, {
        method: "GET",
        headers: {
          "Authorization": "Bearer " + token,
        }
      });

      const data = await res.json();
      loadingIndicator.classList.add("hidden");

      if (!res.ok) {
        if (res.status === 404) {
          noResults.classList.remove("hidden");
        } else {
          errorMessage.classList.remove("hidden");
          errorText.textContent = data.data || "Terjadi kesalahan tidak diketahui.";
        }
        return;
      }

      const user = data.data;
      searchResults.classList.remove("hidden");

      // Buat user card
      const card = document.createElement("div");
      card.className = "bg-gray-500 p-6 rounded-xl border shadow";
      card.innerHTML = `
        <a href="user_profile.html?username=${encodeURIComponent(user.Username)}" class="text-xl font-bold mb-2 text-white hover:underline block">
          @${user.Username}
        </a>
        `;


      userCards.appendChild(card);
    } catch (err) {
      loadingIndicator.classList.add("hidden");
      errorMessage.classList.remove("hidden");
      errorText.textContent = "Gagal terhubung ke server.";
      console.error(err);
    }
  });
});
