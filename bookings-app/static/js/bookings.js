
function Reservation() {
    let prompt = function () {
        let html = `
            <form id="check-availability-form" action="" method="post" class="needs-validation" novalidate>
                <div class="row" id="reservation-dates-modal" style="width: 100%; margin: 0 auto 0 auto;">
                    <div class="col-6">
                        <input name="start" type="text" required class="form-control" id="start"
                            placeholder="Arrival date">
                    </div>
                    <div class="col-6">
                        <input name="end" type="text" required class="form-control" id="end"
                            placeholder="Departure date">
                    </div>
                </div>
            </form>
        `;
        attention.custom({
            title: "Choose your dates",
            msg: html,
        });
    }
    return {
        prompt: prompt,
    };
}

let reserve = Reservation();
