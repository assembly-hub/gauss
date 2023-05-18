package gauss

import (
	"database/sql"
	"fmt"
	"time"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/assembly-hub/db"
	"github.com/assembly-hub/impl-db-sql"
)

type Config struct {
	User            string
	Password        string
	Host            string
	Port            int
	Database        string
	SSLMode         string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifeTime int
	ConnMaxIdleTime int
}

type Client struct {
	cfg *Config
}

func NewClient(cfg *Config) *Client {
	c := new(Client)
	c.cfg = cfg
	return c
}

func (c *Client) Connect() (db.Executor, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		c.cfg.User, c.cfg.Password, c.cfg.Host, c.cfg.Port, c.cfg.Database, c.cfg.SSLMode)
	conn, err := sql.Open("opengauss", dsn)
	if err != nil {
		return nil, err
	}
	conn.SetConnMaxLifetime(time.Duration(c.cfg.ConnMaxLifeTime) * time.Millisecond)
	conn.SetConnMaxIdleTime(time.Duration(c.cfg.ConnMaxIdleTime) * time.Millisecond)
	conn.SetMaxOpenConns(c.cfg.MaxOpenConn)
	conn.SetMaxIdleConns(c.cfg.MaxIdleConn)
	return impl.NewDB(conn), err
}
