import { ConnectInvitation } from "./connectInvitation";
import { Education } from "./education";
import { Experience } from "./experience";
import { Location } from "./location";
import { Message } from "./message";
import { Notification } from "./notification";
import { Post } from "./post";

export type User = {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  additionalName: string;
  profilePhotoUrl: string;
  backgroundPhotoUrl: string;
  headline: string;
  pronouns: string;
  profileLink: string;
  about: string;
  location: Location;
  profileViews: number;
  isActive: boolean;
  experiences: Array<Experience>;
  educations: Array<Education>;
  connections: Array<User>;
  followers: Array<User>;
  following: Array<User>;
  posts: Array<Post>;
  invitations: Array<ConnectInvitation>;
  notifications: Array<Notification>;
  messages: Array<Message>;
  userMightKnow: Array<User>;
  blocked: Array<User>;
};
