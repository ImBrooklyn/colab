package errs

import (
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// common error
var (
    DBError          = status.Error(codes.Internal, "database error")
    RedisError       = status.Error(codes.Internal, "redis error")
    DataParsingError = status.Error(codes.Internal, "data parsing error")
    RPCError         = status.Error(codes.Internal, "rpc error")
)

// colab api error
var (
    TokenExpiry    = status.Error(codes.Unauthenticated, "token is expired")
    TokenNotActive = status.Error(codes.Unauthenticated, "token not active yet")
    TokenMalformed = status.Error(codes.Unauthenticated, "that is not even a token")
    TokenInvalid   = status.Error(codes.Unauthenticated, "could not handle this token")
    NotSignedIn    = status.Error(codes.Unauthenticated, "not signed in")
)

// colab user error
var (
    UsernameExists     = status.Error(codes.AlreadyExists, "username already exists")
    EmailExists        = status.Error(codes.AlreadyExists, "email already exists")
    MobileExists       = status.Error(codes.AlreadyExists, "mobile already exists")
    VerifyCodeNotExist = status.Error(codes.InvalidArgument, "verify code does not exits or expiry")
    VerifyCodeMismatch = status.Error(codes.InvalidArgument, "verify code incorrect")
    SignupFailed       = status.Error(codes.Internal, "signup failed")
    UserNotExist       = status.Error(codes.NotFound, "user does not exist")
    InvalidSigninType  = status.Error(codes.InvalidArgument, "invalid sign type")
    IncorrectPassword  = status.Error(codes.InvalidArgument, "incorrect password")
    MissingUserInfo    = status.Error(codes.InvalidArgument, "missing user info")
)
