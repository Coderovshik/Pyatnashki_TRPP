package ru.solonchev.learnProject.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import ru.solonchev.learnProject.config.UserRequest;
import ru.solonchev.learnProject.config.UserResponse;
import ru.solonchev.learnProject.service.JwtService;

@RestController
@RequestMapping("/api/v1/controller")
@RequiredArgsConstructor
public class AppController {

    private final JwtService jwtService;

    @GetMapping("/users")
    public ResponseEntity<UserResponse> getUser(
            @RequestHeader(name = "Authorization") String bearerToken
    ) {
        HttpHeaders headers = new HttpHeaders();
        headers.set("Access-Control-Allow-Origin", "*");
        headers.set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE");
        headers.set("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token, Origin, Authorization");
        final String token = bearerToken.substring(7);
        final UserResponse userResponse = UserResponse.builder()
                .id(jwtService.extractId(token))
                .email(jwtService.extractUsername(token))
                .build();
        //return new ResponseEntity<>(userResponse, HttpStatus.OK);
        return ResponseEntity.ok()
                .headers(headers)
                .body(userResponse);
    }
}
