$(document).ready(function () {
    initNewRecipeForm();
});

function initNewRecipeForm() {
    var form = $('#newRecipeForm')

    if ($('#newRecipeForm #ingredient').length == 0) {
        addIngredient();
    }
    if ($('#newRecipeForm #step').length == 0) {
        addStep();
    }

    form.form({
        fields: {
            name: {
                identifier: "name",
                rules: [
                    { type: 'empty', prompt: 'Name' },
                ]
            },
            ingredient: {
                identifier: "ingredient",
                rules: [
                    { type: 'empty', prompt: 'Zutat' },
                ]
            },
            amount: {
                identifier: "amount",
                rules: [
                    { type: 'empty', prompt: 'Menge' },
                    { type: 'number', prompt: 'Menge Zahl' }
                ]
            },
            unit: {
                identifier: "unit",
                rules: [
                    { type: 'empty', prompt: 'Einheit' },
                ]
            },
            step: {
                identifier: "step",
                rules: [
                    { type: 'empty', prompt: 'Schritt' },
                ]
            }
        },
        onSuccess: function () {
            $('#modalLoading').modal('show');

            data = $('#newRecipeForm').form('get values');
            recipe = {};
            recipe.Name = data.name;
            recipe.Ingredients = [];
            for (let i = 0; i < data.ingredient.length; i++) {
                recipe.Ingredients.push({
                    "Name": data.ingredient[i],
                    "Amount": data.amount[i],
                    "Unit": data.unit[i],
                    "SortOrder": i
                })
            }
            recipe.Steps = data.step;
            
            $.ajax({
                url: "/recipe",
                data: JSON.stringify(recipe),
                contentType: "application/json",
                method: "PUT",
                processData: false
            }).fail(function(){
                alert("Fehler");
            }).done(function (){
                loadTemplate()
            }).always(function(response){
                $('#modalLoading').modal('hide');
            });

            return false;
        },
        onFailure: function () {
            return false;
        }
    }
    );
}

function moveIngredientUp(elem) {
    ingredient = $(elem).parent();
    prev = $(elem).parent().prev();
    if (prev.is("label")) {
        return false;
    }
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

function removeIngredient(elem) {
    $(elem).parent().remove();
    initNewRecipeForm();
}

function addIngredient(elem) {
    tpl = $('#ingredientTpl').clone().attr('id', 'ingredient').toggle();
    if (elem) {
        $(elem).parent().after(tpl);
        initNewRecipeForm();
    } else {
        $('#ingredients').append(tpl);
    }
}

function moveStepUp(elem) {
    step = $(elem).parent().parent();
    prev = $(elem).parent().parent().prev();
    if (prev.is("label")) {
        return false;
    }
    prev.remove();
    step.after(prev);
    initNewRecipeForm();
}

function moveStepDown(elem) {
    step = $(elem).parent();
    next = $(elem).parent().parent().next();
    next.remove();
    step.before(next);
    initNewRecipeForm();
}

function removeStep(elem) {
    $(elem).parent().parent().remove();
    initNewRecipeForm();
}

function addStep(elem) {
    tpl = $('#stepTpl').clone().attr('id', 'step').toggle();
    if (elem) {
        $(elem).parent().parent().after(tpl);
        initNewRecipeForm();
    } else {
        $('#steps').append(tpl);
    }
}