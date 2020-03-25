$(document).ready(function () {
    $('#mainmenu').dropdown();

    $('#buttonHome').click(buttonHomeClick);
    $('#buttonNew').click(buttonNewClick);
    $('#buttonTags').click(buttonTagsClick);
    $('#buttonPrint').click(buttonPrintClick);

    // buttonHomeClick();
    //routing for "bookmarkable" pages, see RECIPEDB-10
    switch (window.location.hash) {
        case "#recipe":
            buttonTagsClick();
            break;
        case "#new":
            buttonNewClick();
            break;
        default:
            break;
    }
});

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
        console.log("ajax calls completed");
        console.log(template[0]);
        console.log(recipes[0]);
        output = Mustache.render(template[0], recipes[0]);
        console.log(output);
        $('#content').html(output);
        $('#modalLoading').modal('hide');
    });
}

function loadTemplate(name) {
    console.log("load template");
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

function buttonTagsClick() {
    $.get("templates/recipe.html", function (recipeTpl) {
        console.log(recipeTpl);
        $('#content').html(recipeTpl);
    });
}

function buttonPrintClick() {
    $('#modalTest').modal('show');
}