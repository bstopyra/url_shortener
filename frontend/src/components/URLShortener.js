import React, { useState } from "react";
import axios from "axios";
import "./URLShortener.css";
import { Button, Container, Row, Col, Form } from "react-bootstrap";

const endpoint = "http://127.0.0.1:8080";

const URLShortener = () => {
  const [url, setUrl] = useState("");
  const [shortUrl, setShortUrl] = useState("");

  const handleURLChange = (e) => {
    setUrl(e.target.value);
  };

  const handleValidation = (url) => {};

  async function handleButtonOnClick(e) {
    e.preventDefault();

    const response = await axios.post(
      endpoint + "/URL",
      url,
      {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      }
    );
    setShortUrl(response.data);
  }

  return (
    <Container>
      <Row className="justify-content-md-center">
        <Col xs lg="3">
          <Form method="POST">
            <div className="text-center">
              <Form.Label className="m-3">URL Shortener</Form.Label>
              <Form.Control
                name="url"
                type="text"
                placeholder="Your URL"
                onChange={handleURLChange}
              />
            </div>
            <div className="d-grid gap-2">
              <Button variant="primary" size="lg" onClick={handleButtonOnClick}>
                Shorten Your URL
              </Button>
              {shortUrl && <a href={shortUrl}>{shortUrl}</a>}
            </div>
          </Form>
        </Col>
      </Row>
    </Container>
  );
};

export default URLShortener;
