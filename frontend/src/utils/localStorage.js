export function getFromLocalStorage(key) {
  const str = localStorage.getItem(key)
  const data = JSON.parse(str)
  return data
}
export function setInLocalStorage(key, value) {
  const data = JSON.stringify(value)
  localStorage.setItem(key, data)
  window.dispatchEvent(new Event('storage'))
}
