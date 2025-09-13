const loadQuiz = async () => {
  const res = await fetch("/quiz");
  const jsonRes = await res.json();
  document.getElementById("number1").innerText = jsonRes.a;
  document.getElementById("number2").innerText = jsonRes.b;
};

const checkQuiz = async () => {
  const answer = document.getElementById("answer");
  const isCorrect = document.getElementById("isCorrect");
  const A = document.getElementById("number1");
  const B = document.getElementById("number2");
  if (!answer.value || answer.value.length != 2) return;
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
    jsonRes = await res.json();
    answer.value = "";
    if (jsonRes.is_correct == true) {
      isCorrect.innerText = "correct";
      await loadQuiz();
    } else {
      isCorrect.innerText = "wrong try again";
      return;
    }
  } catch (error) {
    console.error("Failed to check ", error);
  }
};

window.onload = () => {
  loadQuiz();
};

document.getElementById("submit-answer").addEventListener("click", checkQuiz);
