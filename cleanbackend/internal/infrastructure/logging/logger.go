package logging

import "go.uber.org/zap"

// NewLogger creates and returns a new zap logger instance.
func NewLogger() (*zap.Logger, error) {
    // Note: In production setup, you'd want to use zap.Config for more control,
    // and possibly have different logging levels for console and file output.
    // Here we're using a basic Production Config for simplicity.
    cfg := zap.NewProductionConfig()
    // Customize the config (e.g. change log level) based on environment or config.
    // e.g., cfg.Level.SetLevel(zap.DebugLevel)

    // Build the logger based on the configuration and return
    logger, err := cfg.Build()
    if err != nil {
        return nil, err
    }
    
    // Set the global logger instance (Optional)
    zap.ReplaceGlobals(logger)
    
    return logger, nil
}