// Global variables
var arrayCarousel = [
    cfg.background + '/1.jpg',
    cfg.background + '/2.jpg',
    cfg.background + '/3.jpg',
    cfg.background + '/4.jpg',
    cfg.background + '/5.jpg',
]
counter = 0;
const setImage = function () {
    document.querySelector(".aside").style.transition = "background-image 0.5s ease-in";
    document.querySelector(".aside").style.backgroundSize = "cover";
    document.querySelector(".aside").style.backgroundRepeat = "no-repeat";
    document.querySelector(".aside").style.backgroundImage = "url(" + arrayCarousel[counter] + ")";
    document.querySelector(".aside").style.transition = "background-image 1.0s ease-out";
    counter = (counter + 1) % arrayCarousel.length;
}
setInterval(setImage, 5000);

