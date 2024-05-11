package ru.solonchev.learnProject.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import ru.solonchev.learnProject.config.UserResponse;
import ru.solonchev.learnProject.service.JwtService;

@RestController
@RequestMapping("/api/v1/controller")
@RequiredArgsConstructor
@CrossOrigin
public class AppController {

    private final JwtService jwtService;

    @GetMapping("/users")
    public ResponseEntity<UserResponse> getUser(
            @RequestHeader(name = "Authorization") String bearerToken
    ) {
        final String token = bearerToken.substring(7);
        final UserResponse userResponse = UserResponse.builder()
                .id(jwtService.extractId(token))
                .email(jwtService.extractUsername(token))
                .build();
        return new ResponseEntity<>(userResponse, HttpStatus.OK);
    }
}
