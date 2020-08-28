package producer

type ClusterConfig struct {
	// Brokers is the list of host/port pairs for a Kafka cluster, e.g. host1:port1,host2:port2,...
	Brokers string

	// SecurityProtocol is the protocol used to communicate with brokers
	SecurityProtocol string

	// SASLMechanism is SASL mechanism to be used for authentication
	SASLMechanism string

	// Username for SASL authentication
	Username string

	// Password for SASL authentication
	Password string

	// SSLCertificateVerification is used for enabling SSL certificate verification
	SSLCertificateVerification bool

	// SSLCertificateAuthorityLocation is the file or directory path to CA certificate(s) for verifying the broker's key
	SSLCertificateAuthorityLocation string
}

