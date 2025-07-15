import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface Blog {
  id: number;
  title: string;
  content: string;
  slug: string;
  created_at: string;
}

@Component({
  selector: 'app-blog-list',
  templateUrl: './blog-list.component.html'
})
export class BlogListComponent implements OnInit {
  blogs: Blog[] = [];

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.http.get<Blog[]>('http://localhost:8080/api/blogs').subscribe({
      next: data => {
        this.blogs = data;
        console.log('Blogs received:', data);
      },
      error: err => {
        console.error('Error fetching blogs:', err);
      }
    });
  }
}
