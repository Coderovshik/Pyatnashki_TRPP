package ru.solonchev.learnProject.service;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.io.Decoders;
import io.jsonwebtoken.security.Keys;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Service;
import ru.solonchev.learnProject.user.User;

import java.security.Key;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.function.Function;

@Service
public class JwtService {

    private static final String SECRET_KEY = "92e17dc8bfac0842d5068f383dd93181e78ee27d3abdb3ed6df848c134b446e5";

    public String extractUsername(String token) {
        return extractClaim(token, (claims) -> claims.get(String.valueOf("email"))).toString();
    }

    public Integer extractId(String token) {
        return extractClaim(token, (claims -> (Integer) (claims.get(String.valueOf("user_id")))));
    }

    public String generateToken(Claims claims) {
        return generateToken(new HashMap<>(), claims);
    }

    public String generateToken(
            Map<String, Object> extraClaims,
            Claims claims
    ) {
        return Jwts
                .builder()
                .setClaims(extraClaims)
                .setClaims(claims)
                .setIssuedAt(new Date(System.currentTimeMillis()))
                .setExpiration(new Date(System.currentTimeMillis() + 1000 * 60 * 24))
                .signWith(getSignInKey(), SignatureAlgorithm.HS256)
                .compact();
    }

    public Claims generateClaims(User user) {
        Claims claims = Jwts.claims();
        claims.put("user_id", user.getId());
        claims.put("email", user.getEmail());
        return claims;
    }

    public boolean isTokenValid(String token, UserDetails userDetails) {
        final String username = extractUsername(token);
        return (username.equals(userDetails.getUsername())) && !isTokenExpired(token);
    }

    private boolean isTokenExpired(String token) {
        return extractExpiration(token).before(new Date());
    }

    private Date extractExpiration(String token) {
        return extractClaim(token, Claims::getExpiration);
    }

    public <T> T extractClaim(String token, Function<Claims, T> claimsResolver) {
        final Claims claims = extractAllClaims(token);
        return claimsResolver.apply(claims);
    }

    private Claims extractAllClaims(String token) {
        return Jwts
                .parserBuilder()
                .setSigningKey(getSignInKey())
                .build()
                .parseClaimsJws(token)
                .getBody();
    }

    private Key getSignInKey() {
        byte[] keyBytes = Decoders.BASE64.decode(SECRET_KEY);
        return Keys.hmacShaKeyFor(keyBytes);
    }
}
