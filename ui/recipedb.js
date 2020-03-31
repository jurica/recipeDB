$(document).ready(function () {
    $('#mainmenu').dropdown();
    
    routie({
        'print': buttonPrintClick,
        'new': buttonNewClick,
        'list': buttonHomeClick,
        'home': buttonHomeClick,
        '': buttonHomeClick,
        'show/:recipeId': function(recipeId) {
            console.log("route matched");
            showRecipe(recipeId);
        }
    });
});

function showRecipe(recipeId){
    $('#modalLoading').modal('show');

    $.when(loadTemplate("recipe"), loadRecipe(recipeId)).done(function (template, recipe) {
        console.log(recipe[0]);
        output = Mustache.render(template[0], recipe[0]);
        $('#content').html(output);
        $('#modalLoading').modal('hide');
    });
}

function buttonNewClick() {
    $('#modalLoading').modal('show');

    loadTemplate("recipeForm").done(function (template) {
        $('#content').html(template);
        $('#modalLoading').modal('hide');
    });
}

function buttonHomeClick() {
    $('#modalLoading').modal('show');

    $.when(loadTemplate("recipeSimple"), loadRecipes()).done(function (template, recipes) {
        // console.log("ajax calls completed");
        // console.log(template[0]);
        // console.log(recipes[0]);
        output = Mustache.render(template[0], recipes[0]);
        // console.log(output);
        $('#content').html(output);
        $('#modalLoading').modal('hide');
    });
}

function loadTemplate(name) {
    console.log("load template: "+name);
    return $.ajax({
        url: "templates/" + name + ".html",
    });
}

function loadRecipes() {
    console.log("load recipes");
    return $.ajax({
        url: "recipe",
    });
}

function loadRecipe(recipeId) {
    console.log("load recipe for id:" + recipeId);
    return $.ajax({
        url: "recipe/"+recipeId,
    });
}

function buttonPrintClick() {
    $('#modalTest').modal('show');
}