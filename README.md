<h1 align="center">
    <img src="https://user-images.githubusercontent.com/82606298/170775456-475ffa71-9cf9-4584-9723-b3917ae0aecc.svg" alt="DebilBot" border="0" height="30px"> 
    DebilBot
</h1>


# GovnoedPromoAPI

![Discord](https://img.shields.io/discord/439778332121497612?style=flat-square&logo=discord&logoColor=white&label=Discord%20Server)
![Website](https://img.shields.io/website?url=https%3A%2F%2Fgovnoed.de%2F&up_message=Visit&down_message=Not%20working&style=flat-square&label=govnoed.de)
![GitHub](https://img.shields.io/website?url=https%3A%2F%2Fgithub.com%2FTiDurak%2Fdebilbot&up_message=tidurak%2Fdebilbot&up_color=black&down_message=removed&style=flat-square&logo=github&label=GitHub)


## Description

This project is a refactored backend for the [govnoed.de](https://govnoed.de/) website.

The API uses an external SQLite database, communicates with the website through HTTP API endpoints, and generates promo codes for [debilbot](https://github.com/tidurak/debilbot).


## Project Status

The project is in its early stages. The basic application structure and configuration are currently in place. The core API functionality, database integration, and promo code generation are under development.


## TODO

- Get api request, then work with database and return api response


## FAQ

Q: Why was the backend rewritten?  
A: As of 23 August 2026, the current [govnoed.de](https://govnoed.de/get_key) backend is written in PHP. PHP was the best choice for my website at the time. However, the project quickly became difficult to scale and turned into an unreadable mess.


Q: Why Go instead of Rust, Django, or Node.js?  
A: The goal was to build a lightweight API because my VPS has very limited resources. The choice ultimately came down to Rust and Go. I chose Go because it keeps the code much simpler with only a small loss in performance. I also had no prior experience with Rust.
