const loadQuiz = async () => {
  const res = await fetch("/quiz");
  const jsonRes = await res.json();
  document.getElementById("number1").innerText = jsonRes.a;
  document.getElementById("number2").innerText = jsonRes.b;
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
  try {
    const res = await fetch("/quiz/check", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        a: A.textContent,
        b: B.textContent,
        answer: answer.value.toLowerCase(),
      }),
    });
    const jsonRes = await res.json();
    answer.value = "";
    if (jsonRes.is_correct == true) {
      updateColor(isCorrect, "green");
      isCorrect.innerText = "correct";
      await loadQuiz();
    } else {
      updateColor(isCorrect, "red");
      isCorrect.innerText = "wrong try again";
      return;
    }
  } catch (error) {
    console.error("Failed to check ", error);
  }
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
