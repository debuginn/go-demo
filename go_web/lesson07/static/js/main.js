(function ($) {
    'use strict';
    
    $(window).on('load',function(){
        var preLoder = $(".preloader");
        preLoder.fadeOut(1000);
    });

    $(document).ready(function(){
        $('.img').slick({
            slidesToShow: 1,
            slidesToScroll: 1,
            fade: true,
            asNavFor: '.txt',
            speed: 2000,
            prevArrow: '<a href="#" class="my-btn prev"><i class="fas fa-long-arrow-alt-left"></i></a>',
            nextArrow: '<a href="#" class="my-btn next"><i class="fas fa-long-arrow-alt-right"></i></a>'
        });
        $('.txt').slick({
            slidesToShow: 1,
            slidesToScroll: 1,
            asNavFor: '.img',
            dots: false,
            centerMode: true,
            focusOnSelect: false,
            autoplay: true,
            autoplaySpeed: 2000,
            fade: false,
            pauseOnFocus: false,
            arrows: false,
            speed: 2000,
            centerPadding: 0
        });
    });

    $(window).on('scroll', function(){
        if($(window).scrollTop() > 100) {
            $('.bottom-header').addClass('fixed-header');
            if ($(window).width() < 992) {
                header.removeClass('fixed-header');
            }
        } else {
            $('.bottom-header').removeClass('fixed-header');
        }
    });

}(jQuery));