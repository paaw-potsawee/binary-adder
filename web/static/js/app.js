let option = "";
const loadQuiz = async () => {
  const A = document.getElementById("number1");
  const B = document.getElementById("number2");
  const optionTag = document.getElementById("option");
  loadingText(A, "00000000");
  loadingText(B, "00000000");
  loadingText(optionTag, "......");
  fetch("/quiz")
    .then((response) => response.json())
    .then((data) => {
      clearLoading(A, data.a);
      clearLoading(B, data.b);
      clearLoading(optionTag, data.option);
      option = data.option;
    })
    .catch((err) => console.error(err));
};

const checkQuiz = async (e) => {
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
  loadingText(isCorrect, ".....");
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
        clearLoading(isCorrect, "correct");
        loadQuiz();
      } else {
        updateColor(isCorrect, "red");
        clearLoading(isCorrect, "wrong! try again");
        return;
      }
    })
    .catch((err) => console.error(err));
};

function updateColor(tag, color) {
  tag.style.color = color;
}

function clearLoading(tag, keyword) {
  tag.classList.remove("loading-font");
  tag.innerText = keyword;
}

function loadingText(tag, keyword) {
  tag.classList.add("loading-font");
  tag.innerText = keyword;
}

window.onload = () => {
  loadQuiz();
};

const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

document
  .getElementById("form-answer")
  .addEventListener("submit", (e) => checkQuiz(e));
