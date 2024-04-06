package ru.solonchev.learnProject.auth;

import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import ru.solonchev.learnProject.service.AuthenticationService;

@RestController
@RequestMapping("/api/v1/auth")
@RequiredArgsConstructor
public class AuthenticationController {

    private final AuthenticationService service;

    @PostMapping("/register")
    public ResponseEntity<AuthenticationResponse> register(
            @RequestBody RegisterRequest request
    ) {
        HttpHeaders headers = new HttpHeaders();
        headers.set("Access-Control-Allow-Origin", "*");
        headers.set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE");
        headers.set("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token, Origin, Authorization");
        service.register(request);
        return ResponseEntity.ok()
                .headers(headers)
                .build();
    }

    @PostMapping("/authenticate")
    public ResponseEntity<AuthenticationResponse> authenticate(
            @RequestBody AuthenticationRequest request
    ) {
        HttpHeaders headers = new HttpHeaders();
        headers.set("Access-Control-Allow-Origin", "*");
        headers.set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE");
        headers.set("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token, Origin, Authorization");
        //return ResponseEntity.ok(service.authenticate(request));
        return ResponseEntity.ok()
                .headers(headers)
                .body(service.authenticate(request));
    }
}
