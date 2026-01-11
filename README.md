# Resume Portfolio Website v3

A dynamic, always-up-to-date portfolio website that automatically syncs with your Google Docs resumes and GitHub projects. Built with Go and HTMX for a modern, performant, server-rendered experience.

## What is this?

This is a personal portfolio website that solves a common problem: keeping your online resume current. Instead of manually updating HTML or markdown files, this website:

1. **Pulls resume content directly from Google Docs** - Write and format your resume in Google Docs like you normally would, and the website automatically displays it
2. **Syncs with your GitHub projects** - Automatically fetches your repositories with their metadata, languages, and topics
3. **Supports multiple resume versions** - Maintain different resumes for different roles (e.g., Frontend, Backend, Full-stack) in separate Google Docs

The website acts as a live view into your resumes and projects, eliminating manual updates and ensuring visitors always see your latest work.

## How it works

### Resume Management

The website connects to Google Drive and reads structured resume documents from a designated folder. Each Google Doc represents a different version of your resume (e.g., "Software Engineer - Frontend", "Software Engineer - Backend").

**The resume parser extracts:**
- Personal information (name, email, contacts, location, title)
- Professional summary
- Skills list
- Work experience (company, role, location, duration, responsibilities)
- Education (degree, institution, duration)
- Personal projects mentioned in the resume

Users can switch between different resume versions using a selector, making it easy to showcase role-specific experience.

### Project Showcase

The website fetches your public GitHub repositories and displays them with:
- Project descriptions
- Programming languages used
- GitHub topics/tags
- Stars and forks count
- Last updated timestamp

Visitors can filter projects by topics (e.g., "frontend", "react", "go") to find relevant work samples.

### Tech Stack

- **Backend**: Go (standard library net/http)
- **Frontend**: HTMX for dynamic interactions without heavy JavaScript
- **APIs**: Google Docs API, Google Drive API, GitHub REST API

## API Endpoints

All API endpoints return JSON responses with CORS headers enabled.

### Resume APIs

#### `GET /api/resumes`

Lists all available resume documents from the configured Google Drive folder.

**Response:**
```json
[
  {
    "id": "1abc123xyz",
    "name": "Software Engineer - Frontend"
  },
  {
    "id": "1def456uvw",
    "name": "Software Engineer - Backend"
  }
]
```

**Use case:** Populate a resume selector dropdown to let users choose which version to view.

---

#### `GET /api/resumes/{id}`

Fetches and parses a specific resume document by its Google Doc ID.

**Path Parameters:**
- `id` - The Google Doc ID (e.g., `1abc123xyz`)

**Response:**
```json
{
  "doc_id": "1abc123xyz",
  "doc_name": "Software Engineer - Frontend",
  "name": "Santiago Gomez",
  "email": "santiago@example.com",
  "contacts": [
    "santiago@example.com",
    "linkedin.com/in/santiagoa58",
    "github.com/santiagoa58"
  ],
  "location": "Austin, TX",
  "title": "Senior Frontend Engineer",
  "summary": "Experienced frontend engineer specializing in React, TypeScript, and modern web applications...",
  "skills": [
    "JavaScript/TypeScript",
    "React",
    "Next.js",
    "Tailwind CSS",
    "Go",
    "Python"
  ],
  "experiences": [
    {
      "company": "Tech Company Inc",
      "location": "Austin, TX",
      "duration": "2022 - Present",
      "role": "Senior Frontend Engineer",
      "responsibilities": [
        "Led development of customer-facing dashboard using React and TypeScript",
        "Improved page load performance by 40% through code splitting and lazy loading",
        "Mentored junior engineers on React best practices"
      ]
    }
  ],
  "educations": [
    {
      "degree": "B.S. Computer Science",
      "institution": "University of Texas",
      "duration": "2016 - 2020"
    }
  ],
  "personal_projects": [
    {
      "name": "Resume Website",
      "description": "Dynamic portfolio website synced with Google Docs"
    }
  ]
}
```

**Use case:** Display the full resume content with proper formatting and sections.

---

### Project APIs

#### `GET /api/projects`

Fetches all public repositories from the configured GitHub account with metadata and language information.

**Query Parameters:**
- `filters` (optional) - Comma-separated list of GitHub topics to filter by
  - Example: `?filters=frontend,react` returns only projects tagged with both "frontend" AND "react"

**Response (without filters):**
```json
[
  {
    "id": 123456789,
    "name": "resume-website-v3",
    "full_name": "santiagoa58/resume-website-v3",
    "owner": {
      "name": "Santiago Gomez",
      "avatar_url": "https://avatars.githubusercontent.com/u/12345",
      "html_url": "https://github.com/santiagoa58",
      "type": "User"
    },
    "private": false,
    "html_url": "https://github.com/santiagoa58/resume-website-v3",
    "description": "Dynamic resume portfolio built with Go and HTMX",
    "fork": false,
    "url": "https://api.github.com/repos/santiagoa58/resume-website-v3",
    "languages_url": "https://api.github.com/repos/santiagoa58/resume-website-v3/languages",
    "homepage": "https://gomezsantiago.com",
    "language": "Go",
    "forks_count": 2,
    "stargazers_count": 15,
    "watchers_count": 15,
    "size": 512,
    "default_branch": "main",
    "topics": ["golang", "htmx", "portfolio", "resume"],
    "archived": false,
    "disabled": false,
    "visibility": "public",
    "pushed_at": "2024-01-15T10:30:00Z",
    "created_at": "2024-01-01T08:00:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "forks": 2,
    "open_issues": 3,
    "watchers": 15,
    "languages": ["Go", "HTML", "CSS"]
  }
]
```

**Response (with filters: `?filters=frontend`):**
```json
[
  {
    "id": 987654321,
    "name": "react-dashboard",
    "description": "Admin dashboard built with React and TypeScript",
    "topics": ["react", "typescript", "frontend", "dashboard"],
    "languages": ["TypeScript", "JavaScript", "CSS"],
    "stargazers_count": 42,
    "html_url": "https://github.com/santiagoa58/react-dashboard",
    ...
  }
]
```

**Use case:**
- Display all projects in a portfolio grid
- Filter projects by technology (e.g., show only frontend projects)
- Sort projects by stars, last updated, or other criteria

---

## User Experience Flow

### Viewing Resumes

1. User visits the website
2. Website displays the default/first resume from Google Drive
3. User can switch between resume versions using a dropdown selector
4. Resume content loads dynamically via HTMX (no page refresh)
5. All sections are properly formatted (experience, education, skills, etc.)

### Browsing Projects

1. User navigates to the projects section
2. All GitHub repositories are displayed with cards showing:
   - Project name and description
   - Languages used (from GitHub language detection)
   - Topics/tags
   - Stars and forks
   - Link to GitHub repo and live demo (if homepage exists)
3. User can filter projects by clicking on topics/tags
4. Projects update dynamically via HTMX

### Dynamic Updates

The website uses HTMX to provide a smooth, app-like experience:
- Resume switching happens without page reload
- Project filtering updates instantly
- All content is server-rendered (better SEO and performance)
- Minimal JavaScript for maximum compatibility

## Data Integrations

### Google Docs Integration

**Required APIs:**
- Google Docs API (for reading document content)
- Google Drive API (for listing documents in folder)

**Authentication:**
- Service account with JSON credentials
- Service account must have read access to the resume folder

**How resume parsing works:**
1. Connect to Google Drive using service account credentials
2. List all documents in the configured folder (`GOOGLE_RESUME_FOLDER_ID`)
3. For each document, fetch the structured content from Google Docs API
4. Parse the document structure (headings, paragraphs, lists) into resume sections
5. Extract data into the resume model (name, experience, skills, etc.)

**Resume document format expectations:**
- Use headings for section titles (e.g., "Experience", "Education", "Skills")
- Use bullet lists for responsibilities and skills
- Follow consistent formatting for dates and locations
- Include contact information at the top

### GitHub Integration

**Required API:**
- GitHub REST API v3

**Authentication:**
- Personal Access Token with `repo` scope (for public repos, `public_repo` is sufficient)

**How project fetching works:**
1. Fetch all repositories for the configured user (`PROJECTS_USERNAME`)
2. For each repository, make a separate API call to get language breakdown
3. Combine repository metadata with language data
4. Filter out forks and archived repos (optional)
5. Apply any topic filters from the query string
6. Return enriched project data

**Rate limiting:**
- GitHub API allows 5,000 requests/hour with authentication
- Each project fetch requires 1 + N requests (1 for list, N for languages)
- Consider caching responses to reduce API calls

## Environment Configuration

The `.env` file contains all necessary API credentials and configuration:

### Google Docs Setup

1. Create a Google Cloud project
2. Enable Google Docs API and Google Drive API
3. Create a service account
4. Download the JSON credentials
5. Extract values into environment variables
6. Share your resume folder with the service account email
7. Copy the folder ID from the Google Drive URL

### GitHub Setup

1. Go to GitHub Settings → Developer settings → Personal access tokens
2. Generate a new token with `repo` or `public_repo` scope
3. Copy the token to `PROJECTS_ACCESS_TOKEN`
4. Set your GitHub username in `PROJECTS_USERNAME`

## Development

```bash
# Install dependencies
go mod download

# Run development server
go run cmd/main.go

# Build for production
go build -o bin/resume-server cmd/main.go
```

## Previous Version

This is a rewrite of the previous Python/React version:
- **Repository**: [resume](https://github.com/santiagoa58/resume)
- **Architecture**: AWS Lambda (Python) + React SPA + S3/CloudFront
- **Live Site**: [gomezsantiago.com](https://gomezsantiago.com)

## What's different in v3?

- **Single binary deployment** instead of serverless functions
- **Server-side rendering** with Go templates instead of client-side React
- **HTMX** for dynamic interactions instead of heavy JavaScript framework
- **Simplified architecture** - one server instead of Lambda + S3 + CloudFront + API Gateway
- **Lower operational costs** - can run on a $5/month VPS
- **Better SEO** - server-rendered HTML instead of SPA

## License

MIT License
