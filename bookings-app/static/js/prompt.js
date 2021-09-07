function Prompt() {
    let notify = function (c) {
        const {
            msg = "",
                msgType = "success"
        } = c;
        notie.alert({
            type: msgType,
            text: msg,
        })
    }

    let notifyModal = function (c) {
        const {
            title = "",
                msg = "",
                icon = "info",
                btnText = "OK",
        } = c;
        Swal.fire({
            title: title,
            html: msg,
            icon: icon,
            confirmButtonText: btnText
        })
    }

    let toast = function (c) {
        const {
            msg = "",
                icon = "success",
                position = "top-end",
        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            icon: icon,
            position: position,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }

    let success = function (c) {
        const {
            msg = "",
                title = "",
                footer = "",
        } = c;

        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    let error = function (c) {
        const {
            msg = "",
                title = "",
                footer = "",
        } = c;

        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    let custom = async function (c) {
        const {
            html = "",
                title = "",
                callback = undefined,
                willOpen = undefined,
                didOpen = undefined,
                preConfirm = undefined,
        } = c;

        Swal.fire({
            title: title,
            html: html,
            focusConfirm: false,
            showCancelButton: true,
            willOpen: () => {
                if (willOpen !== undefined) {
                    willOpen();
                }
            },
            didOpen: () => {
                if (didOpen !== undefined) {
                    didOpen();
                }
            },
            preConfirm: () => {
                if (preConfirm !== undefined) {
                    return preConfirm();
                } else {
                    return {};
                }
            }
        }).then((result) => {
            if (result.isConfirmed && callback !== undefined) {
                if (result.value !== "") {
                    callback(result.value);
                } else {
                    callback(false);
                }
            } else {
                callback(false);
            }
        })
    }

    return {
        notify: notify,
        notifyModal: notifyModal,
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    };
}

let attention = Prompt();