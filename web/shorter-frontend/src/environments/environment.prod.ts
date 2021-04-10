export const environment = {
  production: true,
  apiUrl: window["env"]["apiUrl"] || "http://127.0.0.1:3000",
  debug: window["env"]["debug"] || false
};
