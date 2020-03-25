$(document).ready(function () {
    initNewRecipeForm();
});

function initNewRecipeForm() {
    console.log("test");
    console.log($('#newRecipeForm')
        .form({
            fields: {
                name: {
                    identifier: "dv_ingredient",
                    rules: [
                        { type: 'empty', prompt: 'Please enter numbers' },
                        { type: 'integer', prompt: 'Please enter valid numbers' }
                    ]
                }
            },
            onSuccess: function () {
                form = $('#newRecipeForm').form('get values');
                console.log(form);

                return false;
            },
            onFailure: function () {
                form = $('#newRecipeForm').form('get values');
                console.log(form);

                return false;
            }
        }
        ));
}

function moveIngredientUp(elem) {
    ingredient = $(elem).parent();
    prev = $(elem).parent().prev();
    console.log(prev.is("label"));
    if (prev.is("label")){
        return false;
    }
    console.log(prev);
    prev.remove();
    ingredient.after(prev);
    initNewRecipeForm();
}

function moveIngredientDown(elem) {
    ingredient = $(elem).parent();
    next = $(elem).parent().next();
    next.remove();
    ingredient.before(next);
    initNewRecipeForm();
}

function removeIngredientRow(elem) {
    $(elem).parent().remove();
    initNewRecipeForm();
}

function addIngredientRow(elem) {
    tpl = $('#ingredientTpl').clone().removeAttr('id').toggle();
    $(elem).parent().after(tpl);
    initNewRecipeForm();
}