let attention = Prompt();
const button = document.querySelector("#check-availability-button");
button.addEventListener("click", function (event) {
  let html = `
                <form id="check-availability-form action="" method="post" novalidate
                class = "needs-validation">
                    <div class = "form-row">
                        <div class="col">
                            <div class="form-row" id="reservation-dates-modal">
                                <div class="col">
                                    <input disabled required class = "form-control" type="text" name="start" id="start" placeholder="Arrival">
                                </div>
                                <div class="col">
                                    <input disabled required class = "form-control" type="text" name="end" id="end" placeholder="Departure">
                                </div>
                            </div>
                        </div>
                    </div>
                </form>
                `;
  //notify("You clicked me!", "error");
  //button.classList.toggle("red");
  //notifyModal("title", "<em>hello World</em>", "success", "My text")
  //attention.toast({ msg: "Hello,world!" });
  //attention.success({ msg: "Hello,Wold!" });
  //attention.error({ msg: "error!" });
  attention.custom({
    msg: html,
    title: "choose your dates",
    callback: function (result) {
      fetch("/search-availability-json")
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          console.log(data.ok);
          console.log(data.message);
        });
    },
  });
});
