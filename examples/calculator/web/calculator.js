$(document).ready(function() {
    $("input[name=calc]").click(expressionChanged);
    expressionChanged();
})

function expressionChanged() {
    var operand1 = $("input[name=operand1]").val(),
        operand2 = $("input[name=operand2]").val(),
        operator = $("select[name=operator]").val(),
        resultField = $("input[name=result]"),
        errorField = $("input[name=error]");

    if (operand1 == "") {
        resultField.val("");
        errorField.val("No first operand!");
        return
    }

    op1 = parseInt(operand1);
    if (isNaN(op1)) {
        resultField.val("");
        errorField.val("First Operand is no number!");
        return
    }

    if (operand2 == "") {
        resultField.val("");
        errorField.val("No second operand!");
        return
    }

    op2 = parseInt(operand2);
    if (isNaN(op2)) {
        resultField.val("");
        errorField.val("Second Operand is no number!");
        return
    }

    $.post("/calculate", JSON.stringify({
        operand1: op1,
        operand2: op2,
        operator: operator
    }), function(data) {
        resultField.val(data.result);
        errorField.val("");
    }, 'json').fail(function(error) {
        resultField.val("");
        errorField.val(error.responseText);
    });
}