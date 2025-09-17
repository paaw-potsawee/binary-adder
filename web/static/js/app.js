let option = "";
const loadQuiz = () => {
  const A = document.getElementById("number1");
  const B = document.getElementById("number2");
  const optionTag = document.getElementById("option");
  A.innerText = "loading ...";
  B.innerText = "loading ...";
  optionTag.innerText = "loading ...";
  fetch("/quiz")
    .then((response) => response.json())
    .then((data) => {
      A.innerText = data.a;
      B.innerText = data.b;
      optionTag.innerText = data.option;
      option = data.option;
    })
    .catch((err) => console.error(err));
};

const checkQuiz = (e) => {
  e.preventDefault();
  const answer = document.getElementById("answer");
  const isCorrect = document.getElementById("isCorrect");
  const A = document.getElementById("number1");
  const B = document.getElementById("number2");

  const hexRegex = /^[0-9a-fA-F]{2}$/;
  if (!answer.value || !hexRegex.test(answer.value)) {
    updateColor(isCorrect, "yellow");
    isCorrect.innerText = "Please enter 2-digit hex number";
    answer.value = "";
    return;
  }
  isCorrect.innerText = "loading ...";
  fetch("/quiz/check", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      a: A.textContent,
      b: B.textContent,
      answer: answer.value.toLowerCase(),
      option: option,
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      answer.value = "";
      if (data.is_correct == true) {
        updateColor(isCorrect, "green");
        isCorrect.innerText = "correct";
        loadQuiz();
      } else {
        updateColor(isCorrect, "red");
        isCorrect.innerText = "wrong try again";
        return;
      }
    })
    .catch((err) => console.error(err));
};

function updateColor(tag, color) {
  tag.style.color = color;
}

window.onload = () => {
  loadQuiz();
};

document
  .getElementById("form-answer")
  .addEventListener("submit", (e) => checkQuiz(e));
