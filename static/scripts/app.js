function notify(msg, msgType) {
  notie.alert({
    type: msgType, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
    text: msg,
    stay: false, // optional, default = false
    time: 3, // optional, default = 3, minimum = 1,
    position: "top", // optional, default = 'top', enum: ['top', 'bottom']
  });
}

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

    const { value: result } = await Swal.fire({
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

    if (result) {
      if (result.dismiss !== Swal.DismissReason.cancel) {
        if (result.value !== "") {
          if (c.callback !== undefined) {
            c.callback(result);
          } else {
            c.callback(false);
          }
        } else {
          c.callback(false);
        }
      }
    }
  }
  return {
    toast,
    success,
    error,
    custom,
  };
}
