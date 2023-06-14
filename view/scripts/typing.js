const Q = ["apple", "banana", "melon"];
let Q_No = Math.floor(Math.random() * Q.length);
let Q_i = 0;
let Q_l = Q[Q_No].length;

window.addEventListener("keydown", push_Keydown);

function push_Keydown(event) {
  let keyCode = event.key;
  if (Q_l == Q_l - Q_i) {
    document.getElementById("start").innerHTML = Q[Q_No].substring(Q_i, Q_l);
  }

  if (Q[Q_No].charAt(Q_i) == keyCode) {
    Q_i++;
    document.getElementById("start").innerHTML = Q[Q_No].substring(Q_i, Q_l);
    if (Q_l - Q_i === 0) {
      Q_No = Math.floor(Math.random() * Q.length);
      Q_i = 0;
      Q_l = Q[Q_No].length;
      document.getElementById("start").innerHTML = Q[Q_No].substring(Q_i, Q_l);
    }
  }
}
