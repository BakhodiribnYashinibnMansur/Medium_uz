-- //////////////////////////////////////////////////////////////////////////
-- LIKE TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION like_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_like_count = post_like_count + 1,
    updated_at = NOW()
    WHERE id = NEW.post_id AND deleted_at IS NULL ;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER like_post_trigger AFTER INSERT ON liked_post
FOR EACH ROW EXECUTE PROCEDURE like_post();

-- //////////////////////////////////////////////////////////////////////////
-- UNLIKE TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION unlike_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_like_count = post_like_count - 1,
    updated_at = NOW()
    WHERE id = NEW.post_id AND deleted_at IS NULL ;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER unlike_post_trigger AFTER UPDATE ON liked_post
FOR EACH ROW EXECUTE PROCEDURE unlike_post();


-- //////////////////////////////////////////////////////////////////////////
-- VIEW TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION view_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_views_count = post_views_count + 1,
    updated_at = NOW()
    WHERE id = NEW.post_id AND deleted_at IS NULL ;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER view_post_trigger AFTER INSERT ON viewed_post
FOR EACH ROW EXECUTE PROCEDURE view_post();

-- //////////////////////////////////////////////////////////////////////////
-- RATING TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION rating_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_vote_count = post_vote_count + 1,
    updated_at = NOW()
    WHERE id = NEW.post_id AND deleted_at IS NULL ;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER rating_post_trigger AFTER INSERT ON rating_post
FOR EACH ROW EXECUTE PROCEDURE rating_post();


-- //////////////////////////////////////////////////////////////////////////
--  LIKE FUNCTION
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE  FUNCTION toggle_comment_like(user_id INTEGER,like_post_id INTEGER) RETURNS VOID LANGUAGE PLPGSQL AS
$$
  BEGIN
IF  EXISTS (SELECT id  FROM liked_post  WHERE reader_id =user_id AND post_id=like_post_id AND deleted_at IS NULL)
THEN
UPDATE  liked_post SET deleted_at = NOW() WHERE reader_id = user_id AND post_id = like_post_id  AND deleted_at IS NULL ;
   ELSE
   INSERT INTO liked_post (reader_id  , post_id ) VALUES (user_id  , like_post_id)  ;
   END IF;
  END
$$;

-- //////////////////////////////////////////////////////////////////////////
--  FOLLOWING FUNCTION
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE  FUNCTION toggle_following_user(account_id_func INTEGER,following_id_func INTEGER) RETURNS VOID LANGUAGE PLPGSQL AS
$$
  BEGIN
IF  EXISTS (SELECT id  FROM followings  WHERE account_id =account_id_func AND following_id=following_id_func AND deleted_at IS NULL)
THEN
UPDATE  followings SET deleted_at = NOW() WHERE account_id = account_id_func AND following_id = following_id_func AND deleted_at IS NULL  ;
   ELSE
   INSERT INTO followings (account_id  , following_id ) VALUES (account_id_func  , following_id_func)   ;
   END IF;
  END
$$;

-- //////////////////////////////////////////////////////////////////////////
--  FOLLOWER FUNCTION
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE  FUNCTION toggle_follower_user(account_id_func INTEGER,follower_id_func INTEGER) RETURNS VOID LANGUAGE PLPGSQL AS
$$
  BEGIN
IF  EXISTS (SELECT id  FROM followers  WHERE account_id =account_id_func AND follower_id=follower_id_func AND deleted_at IS NULL)
THEN
UPDATE  followers SET deleted_at = NOW() WHERE account_id = account_id_func AND follower_id = follower_id_func AND deleted_at IS NULL  ;
   ELSE
   INSERT INTO followers (account_id  , follower_id ) VALUES (account_id_func  , follower_id_func)  ;
   END IF;
  END
$$;


-- //////////////////////////////////////////////////////////////////////////
-- FOLLOWING  TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION following_account() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE users SET following_count = following_count + 1,
    updated_at = NOW()
    WHERE id = NEW.account_id AND deleted_at IS NULL ;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER following_account_trigger AFTER INSERT ON followings
FOR EACH ROW EXECUTE PROCEDURE following_account();


-- //////////////////////////////////////////////////////////////////////////
-- UNFOLLOWING  TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION unfollowing_account() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE users SET following_count = following_count - 1,
    updated_at = NOW()
    WHERE id = NEW.account_id AND deleted_at IS NULL ;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER unfollowing_account_trigger AFTER UPDATE ON followings
FOR EACH ROW EXECUTE PROCEDURE unfollowing_account();

-- //////////////////////////////////////////////////////////////////////////
-- FOLLOWER TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION follower_account() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE users SET follower_count = follower_count + 1,
    updated_at = NOW()
    WHERE id = NEW.account_id AND deleted_at IS NULL ;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER follower_account_trigger AFTER INSERT ON followers
FOR EACH ROW EXECUTE PROCEDURE follower_account();


-- //////////////////////////////////////////////////////////////////////////
-- UNFOLLOWER  TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION unfollower_account() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE users SET follower_count = follower_count - 1,
    updated_at = NOW()
    WHERE id = NEW.account_id AND deleted_at IS NULL ;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER unfollower_account_trigger AFTER UPDATE ON followers
FOR EACH ROW EXECUTE PROCEDURE unfollower_account();




-- //////////////////////////////////////////////////////////////////////////
-- ADD RATING FUNCTION
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE  FUNCTION add_rating(user_id INTEGER,rating_post_id INTEGER,reader_rate_func INTEGER ) RETURNS VOID LANGUAGE PLPGSQL AS
$$
  BEGIN
IF  EXISTS (SELECT id  FROM rating_post  WHERE reader_id =user_id AND post_id=rating_post_id AND deleted_at IS NULL)
THEN
UPDATE  rating_post SET reader_rate =  reader_rate_func  WHERE reader_id = user_id AND post_id = rating_post_id AND deleted_at IS NULL   ;
   ELSE
   INSERT INTO rating_post (reader_id  , post_id ,reader_rate ) VALUES (user_id  , rating_post_id , reader_rate_func)  ;
   END IF;
  END
$$;




-- //////////////////////////////////////////////////////////////////////////
-- OVERALL RATING  TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION overall_rating() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
DECLARE
overall_rating double precision :=( SELECT ROUND(AVG(reader_rate ),2) FROM rating_post WHERE deleted_at  IS  NULL AND post_id = NEW.post_id) ;
  BEGIN
    UPDATE post SET post_rated = overall_rating ,
    updated_at = NOW()
    WHERE id = NEW.post_id AND deleted_at IS NULL ;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER overall_rating_trigger AFTER INSERT OR UPDATE ON rating_post
FOR EACH ROW EXECUTE PROCEDURE overall_rating();



-- //////////////////////////////////////////////////////////////////////////
-- ADD RATING FUNCTION
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE  FUNCTION add_rating(user_id INTEGER,rating_post_id INTEGER,reader_rate_func INTEGER ) RETURNS VOID LANGUAGE PLPGSQL AS
$$
  BEGIN
IF  EXISTS (SELECT id  FROM rating_post  WHERE reader_id =user_id AND post_id=rating_post_id AND deleted_at IS NULL)
THEN
UPDATE  rating_post SET reader_rate =  reader_rate_func  WHERE reader_id = user_id AND post_id = rating_post_id AND deleted_at IS NULL   ;
   ELSE
   INSERT INTO rating_post (reader_id  , post_id ,reader_rate ) VALUES (user_id  , rating_post_id , reader_rate_func)  ;
   END IF;
  END
$$;
