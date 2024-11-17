#include <httplib.h>

using namespace httplib;

void handle_hello_world(const Request& req, Response& res) {
    res.set_content("Hello World!", "text/plain");
}

void handle_numbers(const Request& req, Response& res) {
    auto numbers = req.matches[1];
    res.set_content(numbers, "text/plain");
}

void handle_user(const Request& req, Response& res) {
    auto user_id = req.path_params.at("id");
    res.set_content(user_id, "text/plain");
}

void handle_body_header_param(const Request& req, Response& res) {
    if (req.has_header("Content-Length")) {
        auto val = req.get_header_value("Content-Length");
    }
    if (req.has_param("key")) {
        auto val = req.get_param_value("key");
    }
    res.set_content(req.body, "text/plain");
}

void handle_stop(const Request& req, Response& res, Server& svr) {
    svr.stop();
}

int main(void) {
    Server svr;

    svr.Get("/hi", handle_hello_world);
    svr.Get(R"(/numbers/(\d+))", handle_numbers);
    svr.Get("/users/:id", handle_user);
    svr.Get("/body-header-param", handle_body_header_param);

    svr.Get("/stop", [&](const Request& req, Response& res) {
        handle_stop(req, res, svr);
    });

    svr.listen("localhost", 1234);
}
