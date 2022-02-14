document.getElementById("appForm").addEventListener("submit", function (evt) {
  evt.preventDefault();
  const workUrlInputValue = document.getElementById("workUrlInput").value;
  const workUrl = new URL(workUrlInputValue);
  const currentUrl = new URL(location);
  currentUrl.pathname = workUrl.pathname;
  location.assign(currentUrl);
});
