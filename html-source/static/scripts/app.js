function notify(msg, msgType) {
  notie.alert({
    type: msgType, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
    text: msg,
    stay: false, // optional, default = false
    time: 3, // optional, default = 3, minimum = 1,
    position: "top", // optional, default = 'top', enum: ['top', 'bottom']
  });
}

const p = document.querySelector("#paragraph");
p.textContent += "This is being added from JS";
p.classList.add("red");

const button = document.querySelector("#colorButton");

let attention = Prompt();
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
  attention.custom({ msg: html, title: "choose your dates" });
});

const elem = document.getElementById("reservation-date");
const rangepicker = new DateRangePicker(elem, {
  format: "yyyy-mm-dd",
});

(function () {
  "use strict";
  window.addEventListener(
    "load",
    function () {
      // Fetch all the forms we want to apply custom Bootstrap validation styles to
      let forms = document.getElementsByClassName("needs-validation");
      // Loop over them and prevent submission
      let validation = Array.prototype.filter.call(forms, function (form) {
        form.addEventListener(
          "submit",
          function (event) {
            if (form.checkValidity() === false) {
              event.preventDefault();
              event.stopPropagation();
            }
            form.classList.add("was-validated");
          },
          false
        );
      });
    },
    false
  );
})();

function notifyModal(title, html, icon, confirmButtonText) {
  Swal.fire({
    title,
    html,
    icon,
    confirmButtonText,
  });
}

function Prompt() {
  let toast = function (c) {
    const { msg = "", icon = "success", position = "top-end" } = c;
    const Toast = Swal.mixin({
      toast: true,
      title: msg,
      position: position,
      icon: icon,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener("mouseenter", Swal.stopTimer);
        toast.addEventListener("mouseleave", Swal.resumeTimer);
      },
    });

    Toast.fire({});
  };

  let success = function (c) {
    const { msg = "", title = "", footer = "" } = c;
    Swal.fire({
      icon: "success",
      title,
      footer,
      text: msg,
    });
  };
  let error = function (c) {
    const { msg = "", title = "", footer = "" } = c;
    Swal.fire({
      icon: "error",
      title,
      footer,
      text: msg,
    });
  };
  async function custom(c) {
    const { msg = "", title = "" } = c;

    const { value: formValues } = await Swal.fire({
      title,
      html: msg,
      focusConfirm: false,
      showCancelButton: true,
      backdrop: false,
      willOpen: () => {
        const elem = document.getElementById("reservation-dates-modal");
        const rp = new DateRangePicker(elem, {
          format: "yyyy-mm-dd",
          showOnFocus: true,
        });
      },
      preConfirm: () => {
        return [
          document.getElementById("start").value,
          document.getElementById("end").value,
        ];
      },
      didOpen: () => {
        document.getElementById("start").removeAttribute("disabled");
        document.getElementById("end").removeAttribute("disabled");
      },
    });

    if (formValues) {
      Swal.fire(JSON.stringify(formValues));
    }
  }
  return {
    toast,
    success,
    error,
    custom,
  };
}
