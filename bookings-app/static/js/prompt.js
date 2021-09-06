function Prompt() {
    let notify = function(c) {
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
            msg = "",
                title = ""
        } = c;

        const {
            value: formValues
        } = await Swal.fire({
            title: title,
            html: msg,
            focusConfirm: false,
            showCancelButton: true,
            willOpen: () => {
                const elem = document.getElementById('reservation-dates-modal');
                const rp = new DateRangePicker(elem, {format: 'yyyy-MM-dd', showOnFocus: true,});
            },
            didOpen: () => {
                document.getElementById('start').removeAttribute('disabled');
                document.getElementById('end').removeAttribute('disabled');
            },
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            }
        })

        if (formValues) {
            Swal.fire(JSON.stringify(formValues))
        }
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
