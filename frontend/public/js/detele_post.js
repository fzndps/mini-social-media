const deleteButton = postCard.querySelector(".delete-post");
deleteButton.addEventListener("click", () => {
  // Konfirmasi opsional
  if (confirm("Are you sure you want to delete this post?")) {
    postCard.remove();
  }
});