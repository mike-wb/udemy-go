function Reservation() {
    let prompt = function (c) {
        const {
            csrf_token = "",
                room = "",
        } = c;
        const date_form =
            /*html*/
            `
            <form id="check-availability-form" action="" method="post" class="needs-validation" novalidate>
                <div class="row w-100 m-auto" id="reservation-dates-modal">
                    <div class="col-6">
                        <input name="start" type="text" required disabled class="form-control" id="start"
                            placeholder="Arrival date">
                    </div>
                    <div class="col-6">
                        <input name="end" type="text" required disabled class="form-control" id="end"
                            placeholder="Departure date">
                    </div>
                </div>
                <div id="reservation-error" class="mt-1 text-center text-danger"></div>
            </form>
        `;

        attention.custom({
            title: "Choose your dates",
            html: date_form,
            willOpen: () => {
                const elem = document.getElementById('reservation-dates-modal');
                const rp = new DateRangePicker(elem, {
                    format: 'mm-dd-yyyy',
                    minDate: Date.now(),
                    showOnFocus: true,
                    todayHighlight: true,
                    orientation: 'top auto',
                });
                        },
            didOpen: () => {
                document.getElementById('start').removeAttribute('disabled');
                document.getElementById('end').removeAttribute('disabled');
            },
            preConfirm: () => {
                let start = document.getElementById("start");
                let end = document.getElementById("end");
                let error = document.getElementById("reservation-error");
                if (start.value === undefined || start.value === "" || end.value === undefined || end.value === "") {
                    error.innerHTML = "Please select the reservation dates"
                    return false;
                } else {
                    error.innerHTML = "";
                    if (rp) {
                        rp.destroy();
                    }
                    return {
                        "start": start.value,
                        "end": end.value
                    }
                }
            },
            callback: function (result) {
                if (result) {
                    console.log("Callback called with: %j", result);
                    let formData = new FormData();
                    for (var key in result) {
                        formData.append(key, result[key]);
                    }
                    formData.append("csrf_token", csrf_token);
                    formData.append("room", room);

                    fetch('/search-availability-json', {
                            method: "post",
                            body: formData,
                        })
                        .then(response => response.json())
                        .then(data => {
                            console.log(data);
                            console.log(data.ok);
                            console.log(data.message);
                        })
                } else {
                    console.log("Callback called with empty");
                }
            },
        });
    }
    return {
        prompt: prompt,
    };
}

let reserve = Reservation();

//<input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
//<input type="hidden" name="room-id" value="` + room_id + `">