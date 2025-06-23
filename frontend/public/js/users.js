const form = document.getElementById("searchForm");
const input = document.getElementById("searchInput");
const resultDiv = document.getElementById("result");

form.addEventListener("submit", async (e) => {
  e.preventDefault();

  const query = input.value.trim();

  // Validasi input kosong
  if (!query) {
    resultDiv.innerHTML = `
      <p class="text-red-500">Username tidak boleh kosong.</p>
    `;
    return;
  }

  try {
    // Ambil data user dari server
    const response = await fetch("https://localhost/auth/users");

    if (!response.ok) {
      throw new Error("Data gagal diambil.");
    }

    const users = await response.json();

    // Cari user berdasarkan username
    const user = users.find((u) => 
      u.username.toLowerCase() === query.toLowerCase()
    );

    // Tampilkan hasil pencarian
    if (user) {
      resultDiv.innerHTML =
        `<div class="bg-green-100 p-4 rounded text-sm">
          <p><strong>Username:</strong> ${user.username}</p>
        </div>`;
    } else {
      resultDiv.innerHTML = 
      `<p class="text-red-500">User tidak ditemukan.</p>`;
    }
  } catch (err) {
    // Tangani error (seperti koneksi gagal)
    resultDiv.innerHTML = `
      <p class="text-red-500">Terjadi kesalahan: ${err.message}</p>
    `;
  }
});
