package ru.solonchev.learnProject.service;

import lombok.RequiredArgsConstructor;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import ru.solonchev.learnProject.auth.AuthenticationRequest;
import ru.solonchev.learnProject.auth.AuthenticationResponse;
import ru.solonchev.learnProject.auth.RegisterRequest;
import ru.solonchev.learnProject.repository.UserRepository;
import ru.solonchev.learnProject.user.Role;
import ru.solonchev.learnProject.user.User;

@Service
@RequiredArgsConstructor
public class AuthenticationService {

    private final UserRepository repository;
    private final PasswordEncoder passwordEncoder;
    private final JwtService jwtService;
    private final AuthenticationManager manager;

    public AuthenticationResponse register(RegisterRequest request) {
        var user = User.builder()
                .email(request.getEmail())
                .password(passwordEncoder.encode(request.getPassword()))
                .role(Role.USER)
                .build();
        repository.save(user);
        var jwtToken = jwtService.generateToken(jwtService.generateClaims(user));
        return AuthenticationResponse.builder()
                .token(jwtToken)
                .build();
    }

    public AuthenticationResponse authenticate(AuthenticationRequest request) {
        manager.authenticate(
                new UsernamePasswordAuthenticationToken(
                        request.getEmail(),
                        request.getPassword()
                )
        );
        var user = repository.findByEmail(request.getEmail()).orElseThrow();
        var jwtToken = jwtService.generateToken(jwtService.generateClaims(user));
        return AuthenticationResponse.builder()
                .token(jwtToken)
                .build();
    }
}
