# tp-consulting-test
# script db
CREATE TABLE points (
    point_id BIGINT NOT NULL AUTO_INCREMENT,
    thai_id BIGINT NOT NULL,
    balance BIGINT NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (point_id)
);

CREATE TABLE accounts (
    thai_id BIGINT NOT NULL,
    mobile_number VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255) NOT NULL,
    sub_district VARCHAR(100) NOT NULL,
    district VARCHAR(100) NOT NULL,
    province VARCHAR(100) NOT NULL,
    zip_code VARCHAR(10) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (thai_id)
);

CREATE TABLE campaigns (
    campaign_code VARCHAR(50) NOT NULL,  -- Adjust length as needed
    point_action VARCHAR(1) NOT NULL,     -- Assuming point action is a single character (A, D, N)
    campaign_name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    provision TEXT NOT NULL,
    start_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (campaign_code)  -- Assuming campaign_code is unique
);